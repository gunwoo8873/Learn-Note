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
	fmt.Println(u.Name, len(u.Name))
	fmt.Println(u.Class, len(u.Class))
	fmt.Println(u.Userid, len(u.Userid))
	fmt.Println(u.UserPassword, len(u.UserPassword))
	fmt.Println(u.UserEmail, len(u.UserEmail))
}
