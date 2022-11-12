package main

import (
	"crypto/elliptic"
	"crypto/rand"

	"github.com/s3vt/keygen"
)

func main() {
	//RSA keys
	do(&keygen.RSAKey{BitSize: 2048})

	do(&keygen.ECKey{ECCurve: elliptic.P224()})
	do(&keygen.ECKey{ECCurve: elliptic.P256()})
	do(&keygen.ECKey{ECCurve: elliptic.P384()})
	do(&keygen.ECKey{ECCurve: elliptic.P521()})

}

func do(keymaker keygen.Keymaker) {
	keymaker.MakeKeys(rand.Reader)
	keymaker.PrintKeys(true)
}
