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

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_listUsers(t *testing.T) {
	pkg.SetupTest(t)
	defer pkg.ShutdownTest(func() {
		model.ClearTables(&model.User{}, &model.Referral{})
	})

	startAPI()

	_, err := pkg.CreateUser("john", "admin@123", 0)
	require.NoError(t, err)

	url := fmt.Sprintf("http://localhost:%s/user", os.Getenv("PORT"))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := &http.Client{}
	httpRes, err := client.Do(req)
	require.NoError(t, err)
	var res []model.User
	d, _ := ioutil.ReadAll(httpRes.Body)
	err = json.Unmarshal(d, &res)
	require.NoError(t, err)
	assert.Len(t, res, 1)
}
