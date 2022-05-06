package model

// Referral is a referral.
type Referral struct {
	Model
	Code   uint32 `gorm:"uniqueIndex;size:255" json:"code"`
	Active bool   `json:"active"`
	UserID uint   `json:"user_id"`
}

// CreateReferral creates a referral.
func CreateReferral(referral *Referral) error {
	tx := db.Create(referral)
	return tx.Error
}

// GetReferralByCode gets the referral with the given code.
func GetReferralByCode(code uint32) *Referral {
	var referral Referral
	res := db.Where("code = ? ", code).Limit(1).Find(&referral)
	if res.RowsAffected > 0 {
		return &referral
	}
	return nil
}

// GetReferralsByUserID returns the referrals created by the user.
func GetReferralsByUserID(userID uint) []Referral {
	var referrals []Referral
	db.Where("user_id = ? ", userID).Find(&referrals)
	return referrals
}

// UpdateReferral updates a referral.
func UpdateReferral(referral *Referral) error {
	tx := db.Save(referral)
	return tx.Error
}
