package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User struct represents a user in the app
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare calls the methods to validate and format the user
func (u *User) Prepare(step string) error {
	err := u.validate(step)
	if err != nil {
		return err
	}

	u.format()

	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("Name is required and cannot be empty!")
	}

	if u.Nickname == "" {
		return errors.New("Nickname is required and cannot be empty!")
	}

	if u.Email == "" {
		return errors.New("Email is required and cannot be empty!")
	}

	errMail := checkmail.ValidateFormat(u.Email)

	if errMail != nil {
		return errors.New("Email invalid format!")
	}

	if step == "CREATE" && u.Password == "" {
		return errors.New("Password is required and cannot be empty!")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)
}
