package routers

import (
	"fmt"

	"bombay.com/old.thing/controllers"
	ser "bombay.com/old.thing/services"
	tigoweb "bombay.com/old.thing/tigoweb"
	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func handleApi(router *mux.Router) *mux.Router {

	var GetUserMid tigoweb.MiddlewareFunc = func(h *tigoweb.BaseHandler) {
		authBackend := ser.InitJWTAuthenticationBackend()
		req := h.Request
		token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			} else {
				return authBackend.PublicKey, nil
			}
		})

		if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
			claims, _ := token.Claims.(jwt.MapClaims)
			h.SetCtxVal("UserId", claims["sub"])
		}
	}

	var bookmarkHandler = &controllers.BookmarkHandler{}
	bookmarkHandler.Pre(GetUserMid)

	var url = map[string]interface{}{
		"/web/cities":     &controllers.CityHandler{},
		"/web/categories": &controllers.CateHandler{},
		"/web/bookmarks":  bookmarkHandler,
	}

	urlPattern := tigoweb.UrlPattern{}
	urlPattern.Init(router, url)

	apiRouter := mux.NewRouter()
	subRouter := apiRouter.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/posts", controllers.PostList)

	subRouter.HandleFunc("/posts/create", controllers.PostCreate).Methods("POST")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostEdit).Methods("GET")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostUpdate).Methods("PUT")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostDelete).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(negroni.HandlerFunc(RequireTokenAuthentication), negroni.Wrap(apiRouter)))

	return router
}
