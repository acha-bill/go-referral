package api

import (
	"errors"
	"net/http"
	"os"
	"referral/pkg"
	"strconv"
	"time"

	"referral/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	JWTDefaultTTL = 72 * time.Hour
)

var (
	// ErrUnauthorized is when ther user is not authorized.
	ErrUnauthorized = errors.New("unauthorized")
)

type (
	// RegisterReq is the request body for /register endpoint.
	RegisterReq struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,min=8"`
	}

	// RegisterRes is the response for /register endpoint.
	RegisterRes struct {
		ID uint `json:"id"`
	}

	// LoginReq is the request body for /login endpoint.
	LoginReq struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// LoginRes is the response for /login endpoint.
	LoginRes struct {
		Token string `json:"token"`
	}
)

func addAuthRoutes(c *echo.Group) {
	c.POST("/auth/register", register)
	c.POST("/auth/login", login)
}

// @Summary      Register
// @Description  registers a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body RegisterReq true "req"
// @Success      200  {object}  RegisterRes
// @Router       /auth/register [post]
func register(c echo.Context) error {
	var req RegisterReq
	var err error
	if err = c.Bind(&req); err != nil {
		return err
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	code, codeStr := uint32(0), c.QueryParam("code")
	if codeStr != "" {
		code64, err := strconv.ParseUint(codeStr, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		code = uint32(code64)

	}
	user, e := pkg.CreateUser(req.Username, req.Password, code)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusCreated, &RegisterRes{ID: user.ID})
}

// @Summary      Login
// @Description  Logs in a user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login body LoginReq true "req"
// @Success      200  {object}  LoginRes
// @Router       /auth/login [post]
func login(c echo.Context) error {
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.GetUserByUsername(req.Username)
	if user == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "one or more details is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "one or more details is incorrect")
	}

	signedToken, e := GetToken(user.ID)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e)
	}
	return c.JSON(http.StatusOK, &LoginRes{Token: signedToken})
}

func GetToken(userID uint) (string, error) {
	ttl, err := time.ParseDuration(os.Getenv("JWT_TTL"))
	if err != nil {
		ttl = JWTDefaultTTL
	}
	claims := &JwtCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
