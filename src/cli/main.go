package main

import (
	"github.com/joho/godotenv"
	"github.com/devhindo/x/src/cli/auth"
	"os"
	"fmt"
)

func main() {

	// Set Twitter API Key
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	os.Getenv("TWITTER_API_KEY")
	os.Getenv("TWITTER_API_SECRET_KEY")

	// construct url

	auth.Foo()
	fmt.Println(auth.Auth())


}






