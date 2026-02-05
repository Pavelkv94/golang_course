package user

import (
	"fmt"
)

type User struct {
	name        string
	age         int
	phoneNumber string
	isClose     bool
	rating      float64
}

func (u User) Greeting() {
	p := &u
	fmt.Println("User while greeting", p)

	fmt.Println("Hello, I'm ", u.name)
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.age = newAge
	}
}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.phoneNumber = newPhoneNumber
	}
}

func (u *User) CloseAccount() {
	u.isClose = true
}

func (u *User) OpenAccount() {
	u.isClose = false
}

func (u *User) RatingUp() {
	if u.rating+1.0 <= 10.0 {
		u.rating += 1.0
	}
}

func (u *User) RatingDown() {
	if u.rating-1.0 >= 0 {
		u.rating -= 1.0
	}
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}

	if rating > 500 {
		return User{}
	}

	if phoneNumber == "" {
		return User{}
	}

	return User{
		name,
		age,
		phoneNumber,
		isClose,
		rating,
	}

}