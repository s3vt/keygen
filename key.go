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

// KeyPair Generatas RSA Key pair
//
// bitsize : Size of key
//
// reader : A reader for random value, default crypto/rand.Reader.
func KeyPair(bitsize int, reader io.Reader) (*rsa.PrivateKey, error) {
	keyReader := reader
	if keyReader == nil {
		keyReader = rand.Reader
	}

	key, err := rsa.GenerateKey(keyReader, bitsize)

	if err != nil {
		return nil, err
	}

	return key, nil
}

func WriteRSAKeysToFile(filename string, privateKey *rsa.PrivateKey) error {
	pvtPemBlock := &pem.Block{
		Type:  RSAPvtKey,
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}

	pvtFileName := fmt.Sprintf("%s.%s", filename, RSAPvtSuff)

	if err := writeToFile(pvtFileName, pem.EncodeToMemory(pvtPemBlock)); err != nil {
		return err
	}

	pubPemBlock := &pem.Block{
		Type:  RSAPubKey,
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)}
	pubFileName := fmt.Sprintf("%s.%s", filename, RSAPubSuff)

	if err := writeToFile(pubFileName, pem.EncodeToMemory(pubPemBlock)); err != nil {
		return err
	}

	return nil
}

func writeToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0700)
}

func CreateCertificate() {

}

func CertificateSigningRequest(key *rsa.PrivateKey) {
	x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{}, key)
}

func SignKey() { panic("Not Implemented") }

func SaveToFile() { panic("Not Implemented") }

func SaveCertificateToFile() { panic("Not Implemented") }

func SaveKeyToFile() { panic("Not Implemented") }

func PrintKey() { panic("Not Implemented") }

func PrintCertificate() { panic("Not Implemented") }

func Print() { panic("Not Implemented") }
