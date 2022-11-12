package main

import (
	"crypto/rand"

	"github.com/s3vt/keygen"
)

func main() {
	//RSA keys
	rsaKey := &keygen.RSAKey{BitSize: 2048}

	rsaKey.MakeKeys(rand.Reader)

	rsaKey.PrintKeys(true)

	rsaKey.WriteKeysToFile("sapan")

}
