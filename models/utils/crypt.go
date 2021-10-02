package utils

import (
	"crypto/aes"
	"crypto/cipher"
	rand2 "crypto/rand"
	"crypto/sha1"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//Genera una stringa ALFABETICA lunga n
func RandStringRunes(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//Crea una hash per la stringa secret utilizzando la cifratura sha1
func CryptSHA1(secret string) ([]byte, error) {
	h := sha1.New()
	_, err := io.WriteString(h, secret)
	if err != nil {
		return nil, err
	}
	hash := h.Sum(nil)

	return hash, err
}

//Cripta la stringa text usando una key come chiave di cifratura mediante l'algoritmo AES
func EncryptAES(key []byte, text string) ([]byte, error) {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand2.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return ciphertext, nil
}

//Decripta la stringa text usando una key come chiave di cifratura mediante l'algoritmo AES reversato
func DecryptAES(key []byte, ct []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ct) < aes.BlockSize {
		return "", errors.Errorf("stringa troppo corta")
	}
	iv := ct[:aes.BlockSize]
	ct = ct[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ct, ct)

	return string(ct), err
}
