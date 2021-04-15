package middleware

import (
	"BMS-back-end/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if viper.GetBool("jwt.enable") == false {
			return next(c)
		}
		/* try to get JWT from the cookie field */
		cookie, err := c.Cookie("bms")
		if err != nil {
			return c.String(http.StatusUnauthorized, "no cookie named bms was found")
		}

		/* check validity of JWT */
		jwtToken, err := utils.ParseJwt(cookie.Value)
		if err != nil {
			return c.String(http.StatusUnauthorized, "jwt invalid, please login again")
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		uid, _ := strconv.ParseUint(claims["sub"].(string), 10, 64)
		c.Set("uid", uid)

		return next(c)
	}
}
