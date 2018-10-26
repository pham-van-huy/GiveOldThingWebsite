package routers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"old.thing/controllers"
	ser "old.thing/services"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("dist"))))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	router = SocketRouter(router)
	router = SetAuthenticationRoutes(router)
	router = handleApi(router)
	router.HandleFunc("/{home:.*}", controllers.HomeIndex).Methods("GET")

	return router
}

func SocketRouter(router *mux.Router) *mux.Router {
	hub := ser.NewHub()

	go hub.Run()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ser.ServeWs(hub, w, r)
	})
	return router
}

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	router.HandleFunc("/registry", controllers.Registry).Methods("POST")
	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")

	return router
}

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := ser.InitJWTAuthenticationBackend()

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
