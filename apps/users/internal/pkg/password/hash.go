package password

import "golang.org/x/crypto/bcrypt"

type (
	Pass = string
	Hash = string
)

func GenerateHash(pass Pass) (Hash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func VerifyPass(hash Hash, pass Pass) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func VerifyWithPass(p Pass) func(Hash) bool {
	return func(h Hash) bool {
		return VerifyPass(h, p)
	}
}
