package clear

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func delete_user_from_db(license string) int {

	l := License{
		License: license,
	}

	url := "https://x-blush.vercel.app/api/user/delete"

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	req.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	status := resp.StatusCode

	return status

}
