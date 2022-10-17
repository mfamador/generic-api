// Package cryptography handles Cassandra query's "page state" value.
package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Cryptography helps to encrypt and decrypt private data
type Cryptography struct {
	secret string
}

// New returns `Cryptography` type
func New(secret string) Cryptography {
	return Cryptography{
		secret: secret,
	}
}

// EncryptAsString returns an encrypted string version of the given data using a secret key
func (c Cryptography) EncryptAsString(data, secret []byte) (string, error) {
	if secret == nil {
		secret = []byte(c.secret)
	}

	val, _, err := c.encrypt(data, secret, true)
	if err != nil {
		return "", err
	}

	return val, nil
}

func (c Cryptography) encrypt(data, secret []byte, isString bool) (string, []byte, error) { //nolint:gocritic
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", nil, fmt.Errorf("new cipher: %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil, fmt.Errorf("new gcm: %w", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", nil, fmt.Errorf("io read: %w", err)
	}

	bytes := aead.Seal(nonce, nonce, data, nil)

	if isString {
		return base64.URLEncoding.EncodeToString(bytes), nil, nil
	}

	return "", bytes, nil
}

// DecryptString returns a decrypted byte version of the given string data using a secret key
func (c Cryptography) DecryptString(data string, secret []byte) ([]byte, error) {
	if secret == nil {
		secret = []byte(c.secret)
	}

	bytes, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("decode string: %w", err)
	}

	return c.decrypt(bytes, secret)
}

func (c Cryptography) decrypt(data, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, fmt.Errorf("new cipher: %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("new gcm: %w", err)
	}

	size := aead.NonceSize()
	if len(data) < size {
		return nil, fmt.Errorf("nonce size: invalid length")
	}

	nonce, text := data[:size], data[size:]

	res, err := aead.Open(nil, nonce, text, nil)
	if err != nil {
		return nil, fmt.Errorf("aead open: %w", err)
	}

	return res, nil
}
