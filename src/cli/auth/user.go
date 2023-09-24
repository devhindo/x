package auth

import (
	"github.com/devhindo/x/src/cli/utils"
)

type User struct {
	State          string `json:"state"`
	Auth_URL       string `json:"auth_url"`
	Code_verifier  string `json:"code_verifier"`
	Code_challenge string `json:"code_challenge"`
	License        string `json:"license"`
}

func newUser() *User {
	u := new(User)
	u.Lock()
	u.generate_code_verifier()
	u.generate_code_challenge()
	u.generate_state(127)
	u.generate_auth_url()

	return u
}

func (u *User) add_user_to_db() {
	POST("https://x-blush.vercel.app/api/auth/add", *u)
}

func (u *User) open_browser_to_auth_url() {
	utils.OpenBrowser(u.Auth_URL)
}
