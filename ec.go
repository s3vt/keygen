package keygen

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

const (
	ECPvtSuff string = "ecdsa"
	ECPubSuff string = "ecdsa.pub"
	ECPvtKey  string = "EC PRIVATE KEY"
	ECPubKey  string = "EC PUBLIC KEY"
)

type ECKey struct {
	//BitSize of the key to be generated
	BitSize           int
	PrivateKey        *ecdsa.PrivateKey
	PublicKey         *ecdsa.PublicKey
	PrivatKeyPemBlock *pem.Block
	PublicKeyPemBlock *pem.Block

	PrivatKeyPemBlockEncoded []byte
	PublicKeyPemBlockEncoded []byte
}

func (r *ECKey) MakeKeys(reader io.Reader) error {
	if reader == nil {
		reader = rand.Reader
	}

	var err error
	r.PrivateKey, err = ecdsa.GenerateKey(reader, reader)
	if err != nil {
		return fmt.Errorf("could not create keys due to %s ", err.Error())
	}

	r.PublicKey = r.PrivateKey.Public().(*ecdsa.PublicKey)

	r.makeKeyPemBlock()
	return nil
}

func (r *ECKey) PrintKeys(printPrivate bool) error {

	if printPrivate {
		err := encodePemBlock(os.Stdout, r.PrivatKeyPemBlock)
		if err != nil {
			return err
		}
	}
	return encodePemBlock(os.Stdout, r.PublicKeyPemBlock)
}

func (r *ECKey) makeKeyPemBlock() {
	r.makePrivateKeyPemBlock()
	r.makePublicKeyPemBlock()
}

func (r *ECKey) makePrivateKeyPemBlock() {
	r.PrivatKeyPemBlock = &pem.Block{
		Type:  ECPvtKey,
		Bytes: x509.MarshalECPrivateKey(r.PrivateKey),
	}
	r.PrivatKeyPemBlockEncoded = pem.EncodeToMemory(r.PrivatKeyPemBlock)

}

func (r *ECKey) makePublicKeyPemBlock() {
	r.PublicKeyPemBlock = &pem.Block{
		Type:  ECPubKey,
		Bytes: x509.MarshalECPrivateKey(r.PublicKey),
	}
	r.PublicKeyPemBlockEncoded = pem.EncodeToMemory(r.PublicKeyPemBlock)

}

func (r *ECKey) WriteKeysToFile(filename string) error {
	pvtKeyFilename := fmt.Sprintf("%s.%s", filename, ECPvtSuff)
	if err := writeFile(pvtKeyFilename, r.PrivatKeyPemBlockEncoded); err != nil {
		return fmt.Errorf("failed to write private key to file %s due to %s", pvtKeyFilename, err.Error())
	}

	pubKeyFilename := fmt.Sprintf("%s.%s", filename, ECPubSuff)
	if err := writeFile(pubKeyFilename, r.PublicKeyPemBlockEncoded); err != nil {
		return fmt.Errorf("failed to write public key to file %s due to %s", pubKeyFilename, err.Error())
	}
	return nil
}
