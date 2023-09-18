package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	//"encoding/base64"
	"fmt"
	"os"
)

func encrypt() {
	secret_key := secret_key()
	fmt.Println(secret_key)

	information := "access_token"

	nonce := make([]byte, 12)
	cipherBlock, err := aes.NewCipher(secret_key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(information), nil)

	// Save the encrypted information to the system.
	encryptedInformationFile, err := os.OpenFile("encrypted-information.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer encryptedInformationFile.Close()

	_, err = encryptedInformationFile.Write(ciphertext)
	if err != nil {
		panic(err)
	}

	decryptedInformation, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decryptedInformation))

}

func secret_key() []byte {
	// Generate a random sequence of 64 bytes.
	secretKey := make([]byte, 64)
	_, err := rand.Read(secretKey)
	if err != nil {
		panic(err)
	}

	// Encode the secret key to a base64 string.
	//encodedSecretKey := base64.StdEncoding.EncodeToString(secretKey)

	//return encodedSecretKey
	return secretKey
}
