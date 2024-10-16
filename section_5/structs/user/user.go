package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName, 
		lastName: lastName, 
		birthDate: birthdate, 
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}