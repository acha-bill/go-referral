package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "referral/docs"

	_ "github.com/joho/godotenv/autoload"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var e *echo.Echo

var (
	signingKey string
	JWTConfig  middleware.JWTConfig
)

func init() {
	godotenv.Load("./../.env")
	e = echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	signingKey = os.Getenv("JWT_SECRET")
	JWTConfig = middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(signingKey),
		Skipper: func(c echo.Context) bool {
			freeRoutes := []string{"/auth/login", "/auth/register", "/user", "/swagger"}
			for _, route := range freeRoutes {
				if strings.Contains(c.Request().RequestURI, route) {
					return true
				}
			}
			return false
		},
	}

	log.Println("signingKey, ", signingKey)
	log.Printf("%+v\n", JWTConfig)
}

// @title Referral program API
// @version 1.0
// @description referral program API
// @termsOfService http://swagger.io/terms/

// @contact.name Bill
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /
func Start() {
	e.Use(middleware.JWTWithConfig(JWTConfig))
	base := e.Group("")
	addAuthRoutes(base)
	addReferralRoutes(base)
	addUserRoutes(base)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
}

func Shutdown(ctx context.Context) {
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
