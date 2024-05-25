package lock

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
	"io"
	"crypto/sha256"
)

func generateKey() []byte {
    key := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, key); err != nil {
        panic(err)
    }
    return key
}

func encryptString(s string, key []byte) string {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    ciphertext := make([]byte, aes.BlockSize+len(s))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(s))

    return hex.EncodeToString(ciphertext)
}

func decryptString(s string, key []byte) string {
    ciphertext, _ := hex.DecodeString(s)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    if len(ciphertext) < aes.BlockSize {
        panic("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    return string(ciphertext)
}

func EncryptLicense(l string) (string,error) { // using sha256
	hash := sha256.Sum256([]byte(l))
	encyptedLicense := hex.EncodeToString(hash[:])
	return encyptedLicense, nil
}