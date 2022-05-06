package api

import (
	"referral/model"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	// ErrorResponse is a generic error response.
	ErrorResponse struct {
		Error string `json:"error,omitempty"`
	}

	// CustomValidator is a custom validator.
	CustomValidator struct {
		validator *validator.Validate
	}

	// JwtCustomClaims are custom claims extending default ones.
	// See https://github.com/golang-jwt/jwt for more examples
	JwtCustomClaims struct {
		AuthLevel uint `json:"authLevel"`
		UserID    uint `json:"userId"`
		jwt.StandardClaims
	}
)

// Validate validates the interface.
func (c CustomValidator) Validate(i interface{}) error {
	return c.validator.Struct(i)
}

// ClaimsFromContext extracts the claims from the http request context.
func ClaimsFromContext(c echo.Context) *JwtCustomClaims {
	token := c.Get("user").(*jwt.Token)
	return token.Claims.(*JwtCustomClaims)
}

// GetUserFromContext returns the full user from the request context.
func GetUserFromContext(c echo.Context) *model.User {
	claims := ClaimsFromContext(c)
	return model.GetUserByID(claims.UserID)
}
