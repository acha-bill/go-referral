package pkg

import (
	log "github.com/sirupsen/logrus"
	"referral/model"

	"golang.org/x/crypto/bcrypt"
)

const (
	NewUserReward               = 10.0
	ReferrerReward              = 10.0
	ReferralActivationThreshold = 5
)

// CreateUser creates a new user.
func CreateUser(username string, password string, code uint32) (*model.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Username: username,
		Password: string(passwordHash),
	}
	if referralCode := model.GetReferralByCode(code); referralCode != nil {
		user.ReferralID = referralCode.ID
	}

	err = model.CreateUser(user)
	if err != nil {
		return nil, err
	}
	if code == 0 {
		return user, nil
	}

	updatedUser, updateErr := creditNewUser(user, code)
	creditReferrer(code)
	if updateErr != nil {
		log.WithField("creditNewUser", user.ID).WithError(err).Error("failed to credit new user")
		return user, nil
	}
	return updatedUser, err
}

func creditNewUser(user *model.User, code uint32) (*model.User, error) {
	referralCode := model.GetReferralByCode(code)
	if referralCode == nil || !referralCode.Active {
		return user, nil
	}
	user.Balance = NewUserReward
	err := model.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func creditReferrer(code uint32) {
	referralCode := model.GetReferralByCode(code)
	if referralCode == nil || !referralCode.Active {
		return
	}

	user := model.GetUserByID(referralCode.UserID)
	if user == nil {
		return
	}

	users := model.ListByReferral(referralCode.ID)
	if len(users) >= ReferralActivationThreshold {
		user.Balance = ReferrerReward
		if err := model.UpdateUser(user); err != nil {
			log.WithField("creditReferrer", user.ID).WithError(err).Error("failed to credit referrer")
		}

		referralCode.Active = false
		if err := model.UpdateReferral(referralCode); err != nil {
			log.WithField("updateReferral", user.ID).WithError(err).Error("failed to deactivate referral")

		}
	}
}
