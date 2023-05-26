package external_kms

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"gitlab.com/distributed_lab/kit/kv"
)

type Client interface {
	GetKey(keyID string, version string) ([]byte, error)
	Sign(keyID string, version string, message []byte) ([]byte, error)
	Verify(keyID string, version string, message []byte, signature []byte) error
}

const (
	ECDSAKeyType = iota
	ED25519KeyType
)

type PrivateKey struct {
	Type int
	Key  crypto.PrivateKey
}

func NewKeyManagementClient(getter kv.Getter) Client {
	clients := map[string]func(cfg map[string]interface{}) (Client, error){
		"vault":           NewVaultClient,
		"google_kms":      NewGoogleKMSClient,
		"aws_kms":         NewAwsKMSClient,
		"azure_key_vault": NewAzureVaultClient,
		"local":           NewLocalClient,
	}
	var cli Client

	for clientName, newClient := range clients {
		rawCfg, err := getter.GetStringMap(clientName)

		if err != nil {
			panic(err)
		}
		if len(rawCfg) != 0 {
			cli, err = newClient(rawCfg)
			if err != nil {
				panic(err)
			}
			break
		}
	}

	return cli
}

func SignECDSA(priv crypto.PrivateKey, message []byte) ([]byte, error) {
	digest := sha256.New()
	if _, err := digest.Write(message); err != nil {
		return nil, fmt.Errorf("Failed to create digest: %v", err)
	}
	msgHashSum := digest.Sum(nil)

	signature, err := ecdsa.SignASN1(rand.Reader, priv.(*ecdsa.PrivateKey), msgHashSum)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func VerifyECDSA(pub crypto.PublicKey, message []byte, signature []byte) error {
	digest := sha256.New()
	_, err := digest.Write(message)
	if err != nil {
		return fmt.Errorf("Failed to create digest: %v", err)
	}
	msgHashSum := digest.Sum(nil)

	ok := ecdsa.VerifyASN1(pub.(*ecdsa.PublicKey), msgHashSum, signature)
	if !ok {
		return fmt.Errorf("Failed to verify signature")
	}

	return nil
}

func ParseECDSAPrivateKey(key []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(key)

	priv, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ParseECDSAPublicKey(key []byte) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("Failed to decode pem")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub.(*ecdsa.PublicKey), nil
}

func SignED25519(priv crypto.PrivateKey, message []byte) ([]byte, error) {
	digest := sha256.New()
	if _, err := digest.Write(message); err != nil {
		return nil, fmt.Errorf("Failed to create digest: %v", err)
	}
	msgHashSum := digest.Sum(nil)

	signature := ed25519.Sign(*priv.(*ed25519.PrivateKey), msgHashSum)

	return signature, nil
}

func VerifyED25519(pub crypto.PublicKey, message []byte, signature []byte) error {
	digest := sha256.New()
	_, err := digest.Write(message)
	if err != nil {
		return fmt.Errorf("Failed to create digest: %v", err)
	}
	msgHashSum := digest.Sum(nil)

	ok := ed25519.Verify(pub.(ed25519.PublicKey), msgHashSum, signature)
	if !ok {
		return fmt.Errorf("Failed to verify signature")
	}

	return nil
}

func ParseED25519PrivateKey(key []byte) (*ed25519.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("Failed to decode pem")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	edPriv := priv.(ed25519.PrivateKey)
	return &edPriv, nil
}

func ParsePrivateKey(src []byte) (key PrivateKey, err error) {
	block, _ := pem.Decode(src)
	var priv crypto.PrivateKey
	switch block.Type {
	case "EC PRIVATE KEY":
		priv, err = ParseECDSAPrivateKey(src)
		key.Type = ECDSAKeyType
	default:
		priv, err = ParseED25519PrivateKey(src)
		key.Type = ED25519KeyType
	}
	key.Key = priv

	return key, err
}

func Sign(key PrivateKey, message []byte) (signature []byte, err error) {
	switch key.Type {
	case ECDSAKeyType:
		signature, err = SignECDSA(key.Key, message)
	case ED25519KeyType:
		signature, err = SignED25519(key.Key, message)
	}

	return signature, err
}

func Verify(key PrivateKey, message []byte, signature []byte) (err error) {
	switch key.Type {
	case ECDSAKeyType:
		priv := key.Key.(*ecdsa.PrivateKey)
		err = VerifyECDSA(priv.Public(), message, signature)
	case ED25519KeyType:
		priv := key.Key.(*ed25519.PrivateKey)
		err = VerifyED25519(priv.Public(), message, signature)
	}

	return err
}
