package pkg

//
//import (
//	"referral/model"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/require"
//)
//
//func TestNewCode(t *testing.T) {
//	m := make(map[uint32]any)
//	for i := 0; i < 100; i++ {
//		c := newCode()
//		_, ok := m[c]
//		require.False(t, ok)
//		m[c] = struct{}{}
//		time.Sleep(1 * time.Millisecond) // add more randomness
//	}
//}
//
//func TestCreateReferral(t *testing.T) {
//	SetupTest(t)
//	defer ShutdownTest(func() {
//		model.ClearTables(&model.Referral{})
//	})
//	user := &model.User{
//		Model:    model.Model{ID: 1},
//		Username: "u1",
//		Password: testPassword,
//		Balance:  0,
//	}
//	referral, err := CreateReferral(user)
//	require.NoError(t, err)
//	require.Equal(t, uint(1), referral.UserID)
//}
