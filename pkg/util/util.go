package util

import "golang.org/x/crypto/bcrypt"

// HashPassword 密码加密
func HashPassword(password string) (hashPassString string) {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPass)
}

// VerifyPassword 密码验证
func VerifyPassword(hashPass, InputPass string) (equal bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(InputPass))
	if err != nil {
		return false
	}
	return true
}
