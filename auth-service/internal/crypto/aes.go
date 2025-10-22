package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

type (
	Cryptor struct {
		AEAD  cipher.AEAD
		FATAL error
	}
)

func hash(password []byte) []byte {
	sum := sha256.Sum256(password)
	return sum[:]
}

func NewCryptor(password []byte) *Cryptor {
	block, err := aes.NewCipher(hash(password))
	if err != nil {
		return &Cryptor{
			FATAL: err,
		}
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return &Cryptor{
			FATAL: err,
		}
	}

	return &Cryptor{
		AEAD:  aead,
		FATAL: nil,
	}
}

func (crp *Cryptor) Seal(data []byte) string {
	nonce := make([]byte, crp.AEAD.NonceSize())
	rand.Read(nonce)

	cipherText := crp.AEAD.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func (crp *Cryptor) Open(data string) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	nonce := cipherText[:crp.AEAD.NonceSize()]
	text := cipherText[crp.AEAD.NonceSize():]

	open, err := crp.AEAD.Open(nil, nonce, text, nil)
	if err != nil {
		return nil, err
	}

	return open, nil
}
