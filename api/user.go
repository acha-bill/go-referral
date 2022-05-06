package api

import (
	"net/http"
	"referral/model"

	"github.com/labstack/echo/v4"
)

func addUserRoutes(c *echo.Group) {
	c.GET("/user", listUsers)
}

// @Summary      list users
// @Description
// @Tags         User
// @Accept       json
// @Produce      json
// @Router       /user [get]
// @success      200 {array} model.User
func listUsers(c echo.Context) error {
	users := model.ListUsers()
	return c.JSON(http.StatusOK, users)
}
