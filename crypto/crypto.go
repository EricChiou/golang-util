package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"golang.org/x/crypto/sha3"
)

// SHA-3 crypto
func Sha3Encrypt(str string, long int) string {
	switch long {
	case 224:
		byte28 := sha3.Sum224([]byte(str))
		return hex.EncodeToString(byte28[:])

	case 256:
		byte32 := sha3.Sum256([]byte(str))
		return hex.EncodeToString(byte32[:])

	case 384:
		byte48 := sha3.Sum384([]byte(str))
		return hex.EncodeToString(byte48[:])

	case 512:
		byte64 := sha3.Sum512([]byte(str))
		return hex.EncodeToString(byte64[:])

	default:
		byte48 := sha3.Sum384([]byte(str))
		return hex.EncodeToString(byte48[:])
	}
}

// RSA crypto
func LoadPublicKey(filePath string) (*rsa.PublicKey, error) {
	byteBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(byteBuffer)
	if block == nil {
		return nil, errors.New("load public key error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publickey := pub.(*rsa.PublicKey)

	return publickey, nil
}

func LoadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	byteBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(byteBuffer))
	if block == nil {
		return nil, errors.New("load private key error")
	}

	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privatekey, nil
}

func Encrypt(text string, publickey *rsa.PublicKey) (string, error) {
	encryptedTextByteBuffer, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, publickey, []byte(text), nil)
	if err != nil {
		return "", err
	}

	encryptedText := base64.StdEncoding.EncodeToString(encryptedTextByteBuffer)
	return encryptedText, nil
}

func Decrypt(encryptedText string, privatekey *rsa.PrivateKey) (string, error) {
	encryptedTextByteBuffer, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	decryptedText, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, privatekey, encryptedTextByteBuffer, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedText), nil
}
