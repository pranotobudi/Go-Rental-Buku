package router

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type M map[string]interface{}

var APPLICATION_NAME = "Rental Buku App"
var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the_secret_key")

type MyClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  int    `json:"role"`
}

func GenerateJWT(email string, role int) (string, error) {

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Email: email,
		Role:  role,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func IsMember(email string, password string) bool {
	return true
}

// func Auth(c *gin.Context) {
// 	tokenString := c.Request.Header.Get("Authorization")
// 	tokenInJWTStruct, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if jwt.GetSigningMethod("HS256") != token.Method {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return JWT_SIGNATURE_KEY, nil
// 	})

// 	if tokenInJWTStruct != nil && err == nil {
// 		fmt.Println("token verified")
// 		c.JSON(http.StatusOK, gin.H{"message": "authorized"})
// 	} else {
// 		result := gin.H{
// 			"message": "not authorized",
// 			"error":   err.Error(),
// 		}
// 		c.JSON(http.StatusUnauthorized, result)
// 		c.Abort()
// 	}
// }

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenInJWTStruct, err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("jwt_token", tokenInJWTStruct) //save token in the context
		c.Next()
	}
}

//TokenValid check if the Claims Field available or not. it is considered valid.
func TokenValid(r *http.Request) (*jwt.Token, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}
	return token, nil
}
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractTokenFromHeader(r)
	tokenInJWTStruct, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return nil, err
	}
	return tokenInJWTStruct, nil
}
func ExtractTokenFromHeader(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	//should return ["Bearer <token>"]
	if len(strArr) == 2 {
		return strArr[1] //return the token only
	}
	return ""
}
