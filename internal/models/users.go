package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;unique"`
	Username string `gorm:"size:255;unique;not null" json:"username"`
	Password string `gorm:"size:255" json:"password"`
	Phone    string
	PlexID   string
}

func (u *User) BeforeSave() error {
	//turn password into hash
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	//remove spaces in username
	if u.Username != "" {
		u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	}

	return nil
}
