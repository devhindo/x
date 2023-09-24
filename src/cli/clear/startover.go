package clear

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devhindo/x/src/cli/lock"
)

func StartOver() {

	license, err := lock.ReadLicenseKeyFromFile()

	if err != nil {
		fmt.Println("no user logged in")
		return
	}

	delete_user_from_db(license)

	lock.ClearLicenseFile()

	fmt.Println("user deleted successfully")

}

// is there anyway better to pass license?
type License struct {
	License string `json:"license"`
}

func delete_user_from_db(license string) {

	l := License{
		License: license,
	}

	url := "https://x-blush.vercel.app/api/user/delete"

	status := post(url, l)

	if status != 200 {
		fmt.Println("error deleting user from db")
	}
}

func post(url string, l License) int {

    jsonBytes, err := json.Marshal(l)
    if err != nil {
        panic(err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    return resp.StatusCode
}