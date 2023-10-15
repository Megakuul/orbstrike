package crypto

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
)

func EncryptUserKey(id int64, secret string) ([]byte, error) {
	hash:=sha256.Sum256([]byte(secret))
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


func DecryptUserKey(userkey []byte, secret string) (int64, error) {
	hash:=sha256.Sum256([]byte(secret))
	block, err := aes.NewCipher(hash[:])
	if err!=nil {
		return -1, err
	}

	gcm, err := cipher.NewGCM(block)
	if err!=nil {
		return -1, err
	}

	if len(userkey)<gcm.NonceSize()+1 {
		return -1, fmt.Errorf("User key length is too short to process.")
	}
	nonce := userkey[:gcm.NonceSize()]
	key := userkey[gcm.NonceSize():]

	dec, err:=gcm.Open(nil, nonce, key, nil)
	if err!=nil {
		return -1, err
	}
	
	return int64(binary.BigEndian.Uint64(dec)), nil
}
