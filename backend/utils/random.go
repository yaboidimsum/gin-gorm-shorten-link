package utils

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func GenerateCode(n int) string{
	// Membuat wadah array kosong tipe byte dengan panjang n slot
	b:= make([]byte,n)
	for i := range b{
		randomIndex := rand.Intn(len(charset));
		b[i] = charset[randomIndex]
	}
	return string(b)
}