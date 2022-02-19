package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("rahasia"),
	})
}

func CreateToken(userid int, email string, idrole int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userid
	claims["email"] = email
	claims["id_role"] = idrole
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("rahasia"))
}

func GetEmail(e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		if email == "" {
			return email, fmt.Errorf("empty email")
		}
		return email, nil
	}
	return "", fmt.Errorf("invalid user")
}

func GetId(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userid := int(claims["id"].(float64))
		if userid == 0 {
			return userid, fmt.Errorf("invalid id")
		}
		return userid, nil
	}
	return 0, fmt.Errorf("invalid user")
}

func GetIdRole(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		id_role := int(claims["id_role"].(float64))
		if id_role == 0 {
			return id_role, fmt.Errorf("invalid id role")
		}
		return id_role, nil
	}
	return 0, fmt.Errorf("invalid user")
}
