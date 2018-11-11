package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "bombay.com/old.thing/models"
	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"golang.org/x/crypto/bcrypt"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func Login(requestUser *models.User) (int, []byte) {
	authBackend := InitJWTAuthenticationBackend()
	infoUser, isAuth := authBackend.Authenticate(requestUser)

	if isAuth {
		token, err := authBackend.GenerateToken(strconv.Itoa(requestUser.ID))

		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			result := map[string]interface{}{
				"token":    token,
				"infoUser": infoUser,
			}
			response, _ := json.Marshal(result)
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
	authBackend := InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(strconv.Itoa(requestUser.ID))
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := InitJWTAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}

func Registry(user *models.User) bool {
	db := DB_Instance()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashPassword)
	result := db.NewRecord(&user)
	db.Create(user)

	return result
}
