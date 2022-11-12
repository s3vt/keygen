package keygen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

const (
	RSAPvtSuff string = "rsa"
	RSAPubSuff string = "rsa.pub"
	RSAPvtKey  string = "RSA PRIVATE KEY"
	RSAPubKey  string = "RSA PUBLIC KEY"
)

type RSAKey struct {
	//BitSize of the key to be generated
	BitSize           int
	PrivateKey        *rsa.PrivateKey
	PublicKey         *rsa.PublicKey
	PrivatKeyPemBlock *pem.Block
	PublicKeyPemBlock *pem.Block

	PrivatKeyPemBlockEncoded []byte
	PublicKeyPemBlockEncoded []byte
}

func (r *RSAKey) MakeKeys(reader io.Reader) error {
	if reader == nil {
		reader = rand.Reader
	}

	var err error
	r.PrivateKey, err = rsa.GenerateKey(reader, r.BitSize)
	if err != nil {
		return fmt.Errorf("could not create keys due to %s ", err.Error())
	}

	r.PublicKey = r.PrivateKey.Public().(*rsa.PublicKey)

	r.makeKeyPemBlock()
	return nil
}

func (r *RSAKey) PrintKeys(printPrivate bool) error {

	if printPrivate {
		err := encodePemBlock(os.Stdout, r.PrivatKeyPemBlock)
		if err != nil {
			return err
		}
	}
	return encodePemBlock(os.Stdout, r.PublicKeyPemBlock)
}

func (r *RSAKey) makeKeyPemBlock() {
	r.makePrivateKeyPemBlock()
	r.makePublicKeyPemBlock()
}

func (r *RSAKey) makePrivateKeyPemBlock() {
	r.PrivatKeyPemBlock = &pem.Block{
		Type:  RSAPvtKey,
		Bytes: x509.MarshalPKCS1PrivateKey(r.PrivateKey),
	}
	r.PrivatKeyPemBlockEncoded = pem.EncodeToMemory(r.PrivatKeyPemBlock)

}

func (r *RSAKey) makePublicKeyPemBlock() {
	r.PublicKeyPemBlock = &pem.Block{
		Type:  RSAPubKey,
		Bytes: x509.MarshalPKCS1PublicKey(r.PublicKey),
	}
	r.PublicKeyPemBlockEncoded = pem.EncodeToMemory(r.PublicKeyPemBlock)

}

func (r *RSAKey) WriteKeysToFile(filename string) error {
	pvtKeyFilename := fmt.Sprintf("%s.%s", filename, RSAPvtSuff)
	if err := writeFile(pvtKeyFilename, r.PrivatKeyPemBlockEncoded); err != nil {
		return fmt.Errorf("failed to write private key to file %s due to %s", pvtKeyFilename, err.Error())
	}

	pubKeyFilename := fmt.Sprintf("%s.%s", filename, RSAPubSuff)
	if err := writeFile(pubKeyFilename, r.PublicKeyPemBlockEncoded); err != nil {
		return fmt.Errorf("failed to write public key to file %s due to %s", pubKeyFilename, err.Error())
	}
	return nil
}
