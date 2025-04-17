package filecrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	plaintext, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	key := password
	nonce := make([]byte, 12)

	// Randomizing the nonce
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)
	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	// Append the nonce to the end of file
	ciphertext = append(ciphertext, nonce...)

	f, err := os.Create(source)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		return err
	}

	return nil
}

func Decrypt(source string, password []byte) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	ciphertext, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	key := password
	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)
	if err != nil {
		return err
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		return err
	}

	f, err := os.Create(source)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		return err
	}

	return nil
}
