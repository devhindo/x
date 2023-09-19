package auth

import (

)

type User struct {
	State string `json:"state"`
	Auth_URL string `json:"auth_url"`
}

func newUser() *User {
	user := new(User)
	user.Auth_url()
	return user
}

func (u User) send_data_to_server(user User) {
	
}