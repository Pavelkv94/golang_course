package user

import "errors"

type User struct {
	Name    string
	Balance int
}

func (u *User) Pay(usd int) error {
	if u.Balance-usd < 0 {
		return errors.New("not enough money") 
	}
	u.Balance -= usd
	return nil
}
