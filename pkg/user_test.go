package pkg

//
//import (
//	"fmt"
//	"referral/model"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//)
//
//const testPassword = "admin@123"
//
//func TestCreateUser_NoCode(t *testing.T) {
//	SetupTest(t)
//	defer ShutdownTest(func() {
//		model.ClearTables(&model.User{})
//	})
//	user, err := CreateUser("john", testPassword, 0)
//	require.NoError(t, err)
//	assert.Equal(t, "john", user.Username)
//	assert.Equal(t, 0.0, user.Balance)
//}
//
//func TestCreateUser_WithInvalidCode(t *testing.T) {
//	SetupTest(t)
//	defer ShutdownTest(func() {
//		model.ClearTables(&model.User{})
//	})
//	user, err := CreateUser("john", testPassword, 100)
//	require.NoError(t, err)
//	assert.Equal(t, 0.0, user.Balance)
//}
//
//func TestCreateUser_WithValidCode(t *testing.T) {
//	SetupTest(t)
//	defer ShutdownTest(func() {
//		model.ClearTables(&model.User{}, &model.Referral{})
//	})
//	user, err := CreateUser("john", testPassword, 0)
//	require.NoError(t, err)
//	var ref *model.Referral
//	ref, err = CreateReferral(user)
//	require.NoError(t, err)
//
//	// referee gets reward
//	for i := 0; i < ReferralActivationThreshold-1; i++ {
//		u, e := CreateUser(fmt.Sprintf("u%d", i), testPassword, ref.Code)
//		require.NoError(t, e)
//		assert.Equal(t, NewUserReward, u.Balance)
//	}
//
//	// no reward yet
//	user = model.GetUserByID(user.ID)
//	assert.Equal(t, 0.0, user.Balance)
//
//	// referee gets reward
//	u, e := CreateUser(fmt.Sprintf("u%d", ReferralActivationThreshold), testPassword, ref.Code)
//	require.NoError(t, e)
//	assert.Equal(t, NewUserReward, u.Balance)
//
//	// no reward
//	u, e = CreateUser(fmt.Sprintf("u%d", ReferralActivationThreshold+1), testPassword, ref.Code)
//	require.NoError(t, e)
//	assert.Equal(t, 0.0, u.Balance)
//
//	// referer got reward
//	u = model.GetUserByID(user.ID)
//	assert.Equal(t, ReferrerReward, u.Balance)
//}
