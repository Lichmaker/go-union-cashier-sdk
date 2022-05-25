package signaturebuilder

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"os"

	"github.com/pkg/errors"
)

func Sign(bodyString string, privateKeyFilePath string) (string, error) {
	sha256 := crypto.SHA256
	hash256 := sha256.New()
	hash256.Write([]byte(bodyString))
	hashedBody := hash256.Sum(nil)

	privateKeyFile, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		return "", errors.Wrap(err, "load private key file failed")
	}
	pemBlock, _ := pem.Decode(privateKeyFile)
	if pemBlock == nil {
		return "", errors.New("read private key failed")
	}
	priKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return "", errors.New("parse private key failed")
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), sha256, hashedBody)
	if err != nil {
		return "", errors.Wrap(err, "sign failed")
	}

	dst := make([]byte, hex.EncodedLen(len(signature)))
	hex.Encode(dst, signature)
	return string(dst), nil
}
