package people

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const adminUserName = "gocanto"

type AdminUser struct {
	PublicToken  string `validate:"required,min=10"`
	PrivateToken string `validate:"required,min=10"`
}

func (ga AdminUser) IsAllowed(seed string) bool {
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

func (ga AdminUser) IsNotAllowed(seed string) bool {
	return !ga.IsAllowed(seed)
}
