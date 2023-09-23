package clear

import (
	"fmt"
	"net/http"
	"io"
	"bytes"
	"encoding/json"

	"github.com/devhindo/x/src/cli/lock"
)

func StartOver() {

	license, err := lock.ReadLicenseKeyFromFile()

	if err != nil {
		fmt.Println("no license key file")
		return
	}

	status := delete_user_from_db(license)

	if status == 200 {
		fmt.Println("user deleted successfully")
	} else {
		fmt.Println("couldn't delete user")
	}
	

	err = lock.ClearLicenseFile()
	if err != nil {
		fmt.Println("Error clearing license key")
		panic(err)
	} else {
		fmt.Println("Successfully cleared license key")
	}

}

// is there anyway better to pass license?
type License struct {
	License string `json:"license"`
}

func delete_user_from_db(license string) int {

	l := License{
		License: license,
	}

	url := "http://localhost:3000/api/user/delete"

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