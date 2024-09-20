package utils

import (
    "crypto/rand"
    "math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) (string, error) {
    result := make([]byte, length)
    charsetLen := big.NewInt(int64(len(charset)))

    for i := range result {
        randNum, err := rand.Int(rand.Reader, charsetLen)
        if err != nil {
            return "", err
        }
        result[i] = charset[randNum.Int64()]
    }

    return string(result), nil
}
