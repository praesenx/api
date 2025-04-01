package env

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type GlobalAdmin struct {
	PublicToken  string `validate:"required,min=10"`
	PrivateToken string `validate:"required,min=10"`
}

func (ga GlobalAdmin) IsAllowed(seed string) bool {
	token := strings.Trim(ga.PublicToken, " ")
	salt := strings.Trim(ga.PrivateToken, " ")
	externalSalt := strings.Trim(seed, " ")

	if salt != externalSalt {
		return false
	}

	hash := sha256.New()
	hash.Write([]byte(externalSalt))
	bytes := hash.Sum(hash.Sum(nil))

	encodeToString := strings.Trim(
		hex.EncodeToString(bytes),
		" ",
	)

	return token == encodeToString
}

func (ga GlobalAdmin) IsNotAllowed(seed string) bool {
	return !ga.IsAllowed(seed)
}
