package controller

import (
	"BMS-back-end/model"
	"BMS-back-end/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

// @tags Login
// @router /login [post]
// @accept json
// @param data body model.AdminReq true "Login information"
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
