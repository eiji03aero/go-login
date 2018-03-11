package model

import (
	"fmt"
)

type Account struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

func (u *Account) Introduce() {
	fmt.Printf("Name:  %s\n", u.Name)
	fmt.Printf("Password:  %s\n", u.Password)
	fmt.Printf("Profile:\n")
	fmt.Printf("Hi! My name is %s\n", u.Name)
	fmt.Printf("My password is %s\n", u.Password)
	fmt.Printf("%s\n", u.Profile)
}
