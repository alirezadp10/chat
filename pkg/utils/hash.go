package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func Hash(input string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func Verify(input, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(input)) == nil
}
