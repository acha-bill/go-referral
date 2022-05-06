package pkg

import (
	"fmt"
	"hash/fnv"
	"referral/model"
	"time"
)

func newCode() uint32 {
	s := fmt.Sprintf("%d", time.Now().UnixNano())
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// CreateReferral creates a new referral code for the user.
func CreateReferral(user *model.User) (*model.Referral, error) {
	code := newCode()
	referral := &model.Referral{
		Code:   code,
		UserID: user.ID,
		Active: true,
	}
	err := model.CreateReferral(referral)
	return referral, err
}
