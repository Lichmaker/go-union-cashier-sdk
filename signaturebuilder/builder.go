package signaturebuilder

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"os"
	"strings"

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

func Verify(responseData []byte, publicKeyFilePath string) error {
	var getSign struct {
		Sign string `json:"sign"`
	}
	json.Unmarshal(responseData, &getSign)
	if len(getSign.Sign) <= 0 {
		return errors.New("sign is empty")
	}
	signByte, _ := hex.DecodeString(getSign.Sign)

	// 读取证书
	publicKeyFile, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		return errors.Wrap(err, "load pem file failed")
	}
	publicKeyFileString := string(publicKeyFile)
	// 把开头和结尾去掉
	publicKeyFileString = strings.TrimPrefix(publicKeyFileString, "-----BEGIN CERTIFICATE-----\n")
	publicKeyFileString = strings.TrimSuffix(publicKeyFileString, "\n-----END CERTIFICATE-----")
	// 从base64中decode
	certDataDecode, _ := base64.StdEncoding.DecodeString(publicKeyFileString)

	// 解析证书
	certBody, err := x509.ParseCertificate(certDataDecode)
	if err != nil {
		return errors.Wrap(err, "parse pem file failed")
	}
	// // 提取证书中的公钥
	// pubKeyGet, err := x509.MarshalPKIXPublicKey(certBody.PublicKey)
	// if err != nil {
	// 	return errors.Wrap(err, "take public key from PEM file failed")
	// }
	// publicBlock, _ := pem.Decode(pubKeyGet)
	// if publicBlock == nil {
	// 	return errors.New("read public key failed")
	// }
	// pubKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	// if err != nil {
	// 	return errors.Wrap(err, "parse public key failed")
	// }

	// 转成字符串中之后， 找到sign，删掉然后完整保留其他数据
	responseString := string(responseData)
	inx := strings.Index(responseString, "sign")
	new := responseString[0:inx-2] + "}"
	sha256 := crypto.SHA256
	hash256 := sha256.New()
	hash256.Write([]byte(new))
	hashedBody := hash256.Sum(nil)

	err = rsa.VerifyPKCS1v15(certBody.PublicKey.(*rsa.PublicKey), sha256, hashedBody, signByte)

	return err
}
