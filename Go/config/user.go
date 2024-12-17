package config

import "fmt"

// struct to Create Variable and Datatype
type User struct {
	Name         string
	Class        string
	Userid       string
	UserPassword string
	UserEmail    string
}

func NewUser(name string, class string, userid string, userpassword string, useremail string) User {
	return User{
		Name:         name,
		Class:        class,
		Userid:       userid,
		UserPassword: userpassword,
		UserEmail:    useremail,
	}
}

func (u User) Userinfo() {
	fmt.Println(u.Name, u.Class, u.Userid, u.UserPassword, u.UserEmail)
}
