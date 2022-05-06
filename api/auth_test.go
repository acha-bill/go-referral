package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"referral/model"
	"referral/pkg"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Register(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{})
	})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	u := &RegisterReq{
		Username: "john",
		Password: "admin@123",
	}
	d, _ := json.Marshal(u)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(d)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := register(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var res RegisterRes
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	require.NoError(t, err)
	assert.Equal(t, uint(1), res.ID)
}

func Test_login(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{})
	})

	_, err := pkg.CreateUser("john", "admin@123", 0)
	require.NoError(t, err)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	u := &LoginReq{
		Username: "john",
		Password: "admin@123",
	}
	d, _ := json.Marshal(u)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(d)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = login(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var res LoginRes
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	require.NoError(t, err)
	assert.NotEqual(t, "", res.Token)
}
