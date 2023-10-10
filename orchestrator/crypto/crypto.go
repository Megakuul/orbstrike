package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/binary"
	"math/rand"
)

func EncryptUserKey(id int64, secret string) ([]byte, error) {
	hash:=sha512.Sum512([]byte(secret))
	block, err := aes.NewCipher(hash[:])
	if err!=nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err!=nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err!=nil {
		return nil, err
	}

	bId := make([]byte, 8)
	binary.BigEndian.PutUint64(bId, uint64(id))

	enc:=gcm.Seal(nil, nonce, bId, nil)
	return append(nonce, enc...), nil
}
