package keygen

import (
	"crypto/ecdsa"
	"crypto/elliptic"
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
	ECCurve           elliptic.Curve
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
	r.PrivateKey, err = ecdsa.GenerateKey(r.ECCurve, reader)
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

func (r *ECKey) makePrivateKeyPemBlock() error {

	bytes, err := x509.MarshalECPrivateKey(r.PrivateKey)
	if err != nil {
		return fmt.Errorf("could not marshal private key due to %s", err.Error())
	}

	r.PrivatKeyPemBlock = &pem.Block{
		Type:  ECPvtKey,
		Bytes: bytes,
	}
	r.PrivatKeyPemBlockEncoded = pem.EncodeToMemory(r.PrivatKeyPemBlock)
	return nil

}

func (r *ECKey) makePublicKeyPemBlock() error {

	bytes, err := x509.MarshalPKIXPublicKey(r.PublicKey)
	if err != nil {
		return fmt.Errorf("could not marshal public key due to %s", err.Error())
	}

	r.PublicKeyPemBlock = &pem.Block{
		Type:  ECPubKey,
		Bytes: bytes,
	}
	r.PublicKeyPemBlockEncoded = pem.EncodeToMemory(r.PublicKeyPemBlock)

	return nil
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
