package model

import (
	"errors"
)

var (
	ErrDuplicatedUsername = errors.New("username already exists")
)

// User is a user.
type User struct {
	Model
	Username   string     `gorm:"uniqueIndex;size:255" json:"username"`
	Password   string     `json:"-"`
	Balance    float64    `json:"balance"`
	Referrals  []Referral `json:"-"` // referrals created
	ReferralID uint
	Referral   Referral `json:"-"` // referral signed up with
}

// GetUserByID gets the user with the ID specified.
func GetUserByID(ID uint) *User {
	var user User
	res := db.First(&user, ID)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}

// GetUserByUsername gets the user with the username specified.
func GetUserByUsername(username string) *User {
	var user User
	res := db.First(&user, "username = ?", username)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}

// CreateUser inserts a new user in the DB.
func CreateUser(user *User) error {
	var existingUser User
	tx := db.Where("username = ? ", user.Username).Limit(1).Find(&existingUser)
	if tx.RowsAffected > 0 {
		return ErrDuplicatedUsername
	}

	tx = db.Create(user)
	return tx.Error
}

func DeleteUser(user *User) {
	db.Delete(user)
}

// ListUsers list users.
func ListUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

// UpdateUser updates the user.
func UpdateUser(user *User) error {
	tx := db.Save(user)
	return tx.Error
}

// ListByReferral returns the users who signed up with the referral.
func ListByReferral(referralID uint) []User {
	var users []User
	db.Where("referral_id = ?", referralID).Find(&users)
	return users
}
