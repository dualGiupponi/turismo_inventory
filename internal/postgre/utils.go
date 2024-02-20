package postgre

import (
	"crypto/aes"
	"math/rand"
	"time"
)

func generateSeed() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numberRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 20)
	for i := range b {
		b[i] = charset[numberRand.Intn(len(charset))]
	}
	return string(b)
}

func encryptPassword(password string) (string, string) {
	seed := generateSeed()

	encryptor, _ := aes.NewCipher([]byte(seed))

	ciphertext := make([]byte, len(password))
	encryptor.Encrypt(ciphertext, []byte(password))

	encryptPwd := string(ciphertext)
	return encryptPwd, seed
}
