package auth

// func check_authentication() {}

func Auth() {
	user := newUser()
	user.add_user_to_db()
	user.open_browser_to_auth_url()
}

