package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                  string `gorm:"type:varchar(255);not null"`
	Password              string `gorm:"not null"`
	Email                 string `gorm:"unique;not null"`
	EmailVerificationCode string
	EmailVerified         bool `gorm:"not null"`
	IsAllowByAdmin        bool `gorm:"not null"`
}

type UserResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// Create a user
func CreateUser(User *User) (err error) {
	err = Db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all users
func GetUsers(User *[]User) (err error) {
	err = Db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by id
func GetUser(User *User, id int) (err error) {
	err = Db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Update user
func UpdateUser(User *User) (err error) {
	Db.Save(User)
	return nil
}

// Delete user
func DeleteUser(User *User, id int) (err error) {
	Db.Where("id = ?", id).Delete(User)
	return nil
}

// Get user by email verification code
func GetUserByEmailVerifyCode(User *User, verificationCode string) (err error) {
	err = Db.First(&User, "email_verification_code = ?", verificationCode).Error
	if err != nil {
		return err
	}

	return nil
}

// Get user by email
func GetUserByEmail(User *User, email string) (err error) {
	err = Db.First(&User, "email = ?", email).Error
	if err != nil {
		return err
	}

	return nil
}

// Generate and set password hash
func GenerateAndSetPasswordHash(user *User) (err error) {
	// Generate password hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)

	return nil
}
