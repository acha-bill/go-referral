package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"referral/model"
	"referral/pkg"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func Test_createSignupReferral(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{}, &model.Referral{})
	})
	startAPI()

	user, err := pkg.CreateUser("john", "admin@123", 0)
	require.NoError(t, err)
	jwt, err := GetToken(user.ID)
	require.NoError(t, err)

	url := fmt.Sprintf("http://localhost:%s/referral", os.Getenv("PORT"))
	req, err := http.NewRequest(http.MethodPost, url, nil)
	require.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", jwt))

	client := &http.Client{}
	httpRes, err := client.Do(req)
	require.NoError(t, err)
	var res CreateReferralRes
	d, _ := ioutil.ReadAll(httpRes.Body)
	err = json.Unmarshal(d, &res)
	require.NoError(t, err)
	fmt.Println(res.Link)

	assert.Regexp(t, `http:\/\/localhost:\d+\/auth\/signup\?code=\d+`, res.Link)
}

func Test_listReferrals(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{}, &model.Referral{})
	})

	startAPI()

	user, err := pkg.CreateUser("john", "admin@123", 0)
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		_, err = pkg.CreateReferral(user)
		require.NoError(t, err)
	}
	jwt, err := GetToken(user.ID)
	require.NoError(t, err)

	url := fmt.Sprintf("http://localhost:%s/referral", os.Getenv("PORT"))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", jwt))

	client := &http.Client{}
	httpRes, err := client.Do(req)
	require.NoError(t, err)
	var res []model.Referral
	d, _ := ioutil.ReadAll(httpRes.Body)
	err = json.Unmarshal(d, &res)
	require.NoError(t, err)

	assert.Len(t, res, 5)
}
