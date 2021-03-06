package controller

import (
	"BMS-back-end/model"
	"BMS-back-end/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// @tags Login
// @Summary Admin login
// @router /login [post]
// @accept json
// @param data body model.AdminReq true "Login information"
// @produce plain
// @success 200
func login(c echo.Context) error {
	adminReq := model.AdminReq{}
	if err := c.Bind(&adminReq); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	} else if err = c.Validate(&adminReq); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	adminDb := &model.Admin{Name: adminReq.Name}
	if err := model.RetrieveAdmin(adminDb); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusForbidden, "invalid username or password")
		}
		logrus.Error("Retrieve admin failed")
		return c.NoContent(http.StatusInternalServerError)
	}
	if adminReq.Password != adminDb.Password {
		return c.String(http.StatusForbidden, "invalid username or password")
	}

	jwtString, expTime, err := utils.GenerateJwt(adminDb.ID)
	if err != nil {
		logrus.Error("Generate JWT failed")
		return c.NoContent(http.StatusInternalServerError)
	}

	cookie := new(http.Cookie)
	cookie.Name = "bms"
	cookie.Value = jwtString
	cookie.Expires = *expTime
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// @tags Login
// @Summary Check login status
// @router /login [get]
// @success 200 {object} model.AdminResp
// @failure 401 {string} string "Not logged in"
func getLoginStatus(c echo.Context) error {
	if !viper.GetBool("jwt.enable") {
		return c.JSON(http.StatusOK, &model.AdminResp{Name: "Test Account", Contact: "me@example.com"})
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

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		logrus.Error("error while retrieving user data: JWT format error")
		return c.JSON(http.StatusOK, &model.AdminResp{Name: "Test Account", Contact: "me@example.com"})
	}
	uid, err := strconv.ParseUint(claims["sub"].(string), 10, 64)
	if err != nil {
		logrus.WithField("uid", claims["sub"].(string)).
			Error("error while retrieving user data: SUB in JWT is not an integer")
		return c.JSON(http.StatusOK, &model.AdminResp{Name: "Test Account", Contact: "me@example.com"})
	}

	dbAdmin := model.Admin{ID: uid}
	err = model.RetrieveAdmin(&dbAdmin)
	if err != nil {
		logrus.WithField("uid", uid).
			Error("error while retrieving user data: uid not exist in database")
		return c.JSON(http.StatusOK, &model.AdminResp{Name: "Test Account", Contact: "me@example.com"})
	}

	return c.JSON(http.StatusOK, &model.AdminResp{Name: dbAdmin.Name, Contact: dbAdmin.Contact})
}
