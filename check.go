package passwd

import "golang.org/x/crypto/bcrypt"

func Check(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
