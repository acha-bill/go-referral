package api

import (
	"fmt"
	"net/http"
	"referral/model"
	"referral/pkg"

	"github.com/labstack/echo/v4"
)

type (
	// CreateReferralRes is the response for create referral.
	CreateReferralRes struct {
		Link string `json:"link"`
	}
)

func addReferralRoutes(c *echo.Group) {
	c.GET("/referral", listUserReferrals)
	c.POST("/referral", createSignupReferral)
}

// @Summary      Create signup referral
// @Description
// @Tags         Referral
// @Accept       json
// @Produce      json
// @Router       /referral [post]
func createSignupReferral(c echo.Context) error {
	user := GetUserFromContext(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrUnauthorized)
	}
	referral, err := pkg.CreateReferral(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	link := fmt.Sprintf("%s://%s/auth/signup?code=%d", c.Scheme(), c.Request().Host, referral.Code)
	return c.JSON(http.StatusCreated, CreateReferralRes{Link: link})
}

// @Summary      list user's referrals
// @Description
// @Tags         Referral
// @Accept       json
// @Produce      json
// @Router       /referral [get]
// @success      200 {array} model.Referral
func listUserReferrals(c echo.Context) error {
	user := GetUserFromContext(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrUnauthorized)
	}
	referrals := model.GetReferralsByUserID(user.ID)
	return c.JSON(http.StatusCreated, referrals)
}
