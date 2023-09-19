package main

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

func Generate_code_challenge(s string) string {
	// Base64-URL-encoded string of the SHA256 hash of the code verifier
	return base64.RawURLEncoding.EncodeToString([]byte(s))
}

func hash_sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func generate_code_verifier() string {
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

	return string(b)
}