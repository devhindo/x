package main

import (
	//"github.com/devhindo/x/src/cli/x"
	"github.com/devhindo/x/src/cli/api"
)

func main() {
	//x.Run()
	api.GET("http://localhost:3000/api/auth")
}






