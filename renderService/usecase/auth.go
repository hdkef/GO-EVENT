package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var SECRET string
var EXPIRED int64 = int64(10000000) //in second

func init() {
	_ = godotenv.Load()
	SECRET = os.Getenv("SECRET")
}

func SetAuthCookie(c *gin.Context, userID *uint32) error {
	signedToken, err := generateJWT(userID)
	if err != nil {
		return err
	}

	c.Set("AUTH", signedToken)
	return nil
}

func generateJWT(userID *uint32) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        userID,
		"timestamp": time.Now().Unix(),
	})
	signedToken, err := newToken.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateAuthCookie(c *gin.Context) (*uint32, error) {
	signedToken, err := c.Cookie("AUTH")
	if err != nil {
		return nil, err
	}
	id, err := validateToken(&signedToken)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func validateToken(tokenStr *string) (uint32, error) {
	token, err := jwt.Parse(*tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}
	jwtMapClaims := token.Claims.(jwt.MapClaims)
	timestamp := int64(jwtMapClaims["timestamp"].(float64))

	err = isExpired(&timestamp)
	if err != nil {
		return 0, err
	}
	//calculate if jwt is expired

	//get id from jwt claims

	idFloat := jwtMapClaims["ID"].(float64)

	return uint32(idFloat), nil
}

func isExpired(timestamp *int64) error {
	if time.Now().Unix()-*timestamp > EXPIRED {
		return errors.New("token is expired")
	}
	return nil
}
