package auth

type User struct {
	State string `json:"state"`
	Auth_URL string `json:"auth_url"`
	Code_verifier string `json:"code_verifier"`
	Code_challenge string `json:"code_challenge"`
}

func newUser() *User {
	user := new(User)
	user.generate_code_verifier()
	user.generate_code_challenge()
	user.generate_state(127)
	user.generate_auth_url()
	
	return user
}

func (u *User) add_user_to_db() {
	POST("http://localhost:3000/api/auth/add", *u)
}