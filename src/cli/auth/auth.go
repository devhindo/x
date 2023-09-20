package auth

import (
	"fmt"
)

// func check_authentication() {}

func Auth() {
	user := newUser()
	fmt.Println(user)
}

