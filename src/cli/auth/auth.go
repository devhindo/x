package auth

import (
	"github.com/devhindo/x/src/cli/helpers"
	"fmt"
)

func Auth() {
	if(!IsAuthenticated()) {
		url := Auth_url()
		helpers.Open(url)


		fmt.Println("Authenticated")
	} else {
		fmt.Println("Already authenticated")
	}
	fmt.Println("Error")
}

