package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"encoding/base64"
)

/*
*	code_challenge = BASE64URL-ENCODE(SHA256(ASCII(code_verifier)))
*/

func (u *User) generate_code_challenge() {
	// Base64-URL-encoded string of the SHA256 hash of the code verifier
	u.Code_challenge = base64.RawURLEncoding.EncodeToString([]byte(hash_sha256(u.Code_verifier)))
}

func hash_sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (u *User) generate_code_verifier() {
	const (
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~"
		min   = 43
		max   = 128
	)

	length, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		panic(err)
	}
	length.Add(length, big.NewInt(min))

	b := make([]byte, length.Int64())
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		b[i] = chars[n.Int64()]
	}
	
	u.Code_verifier = string(b)
}

func (u *User) generate_state(stateLength int) {
	b := make([]byte, stateLength)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	state := base64.URLEncoding.EncodeToString(b)

	u.State = state
}