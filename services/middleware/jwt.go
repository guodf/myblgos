package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type MyClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// JwtAuth middlerware
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtStr := c.Request.Header.Get("Authorization")
		if jwtStr == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		myClaims := &MyClaims{}
		myToken, err := jwt.ParseWithClaims(jwtStr, myClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt_key")), nil
		})
		log.Println(myToken.Valid, myToken.Claims, myClaims)
		if err != nil || !myToken.Valid {
			var validError *jwt.ValidationError
			if errors.As(err, &validError) {
				if validError.Errors == jwt.ValidationErrorExpired {
					log.Println("授权过期")
				}
			}

			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("claims", myClaims)
		c.Next()
	}
}

// createJwtToken
func CreateJwtToken(name string) (string, int64) {
	exp := time.Now().Add(time.Hour * 24 * 30).Unix()

	myClaims := &MyClaims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenStr, _ := token.SignedString([]byte(viper.GetString("jwt_key")))
	return tokenStr, exp
}
