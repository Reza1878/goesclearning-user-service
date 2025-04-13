package middlewares

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

var salt = []byte("secret-password")

func GenerateHashed(password string) string {
	timeCost := uint32(1)
	memCost := uint32(64 * 1024)
	parallelism := uint8(1)
	hashLength := uint32(32)

	hash := argon2.IDKey([]byte(password), salt, timeCost, memCost, parallelism, hashLength)

	return base64.RawStdEncoding.EncodeToString(hash)
}

func VerifyPassword(hash, password string) bool {
	return hash == string(argon2.IDKey([]byte(password), salt, 1, 64*1024, 1, 32))
}
