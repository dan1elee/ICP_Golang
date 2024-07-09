package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func encryptString(plainString string) [32]byte {
	return sha256.Sum256([]byte(plainString))
}

func EncryptPassword(plainPassword string) string {
	byte_arr := encryptString(plainPassword)
	return hex.EncodeToString(byte_arr[:])
}
