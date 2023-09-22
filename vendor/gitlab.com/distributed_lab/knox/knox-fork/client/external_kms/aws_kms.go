package external_kms

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"gitlab.com/distributed_lab/figure"
	"sync"
)

type AwsKMSClient struct {
	mutex   *sync.RWMutex
	Session *session.Session
	Config  *AwsKMSClientConfig
	keys    map[string]map[string][]byte
}

type AwsKMSClientConfig struct {
	Region      string `fig:"region"`
	AccessKeyID string `fig:"access_key_id"`
	SecretKey   string `fig:"secret_access_key"`
	Token       string `fig:"session_token"`
}

func NewAwsKMSClient(rawCfg map[string]interface{}) (Client, error) {
	config := AwsKMSClientConfig{}
	err := figure.Out(&config).From(rawCfg).Please()
	if err != nil {
		return nil, err
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretKey, config.Token),
	},
	)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]map[string][]byte)
	var mutex = &sync.RWMutex{}
	return &AwsKMSClient{
		mutex:   mutex,
		Session: sess,
		Config:  &config,
		keys:    keys,
	}, nil
}

func (a *AwsKMSClient) getCacheKey(keyID string, version string) []byte {
	a.mutex.RLock()
	cachedKey := a.keys[keyID][version]
	a.mutex.RUnlock()
	return cachedKey
}

func (a *AwsKMSClient) setCacheKey(keyID string, version string, key []byte) {
	a.mutex.Lock()
	a.keys[keyID] = make(map[string][]byte)
	a.keys[keyID][version] = key
	a.mutex.Unlock()
}

func (a *AwsKMSClient) GetKey(keyID string, version string) ([]byte, error) {
	cacheKey := a.getCacheKey(keyID, version)
	if cacheKey != nil {
		return cacheKey, nil
	}

	svc := kms.New(a.Session)
	input := &kms.GetPublicKeyInput{
		KeyId: aws.String(keyID),
	}
	result, err := svc.GetPublicKey(input)
	if err != nil {
		return nil, err
	}

	pub, err := x509.ParsePKIXPublicKey(result.PublicKey)
	if err != nil {
		return nil, err
	}
	publicKeyDer, _ := x509.MarshalPKIXPublicKey(pub)
	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}
	publicKeyPem := pem.EncodeToMemory(&publicKeyBlock)

	a.setCacheKey(keyID, version, publicKeyPem)
	return publicKeyPem, nil
}

func (a *AwsKMSClient) Sign(keyID string, version string, message []byte) ([]byte, error) {
	svc := kms.New(a.Session)
	input := &kms.SignInput{
		KeyId:            aws.String(keyID),
		Message:          message,
		MessageType:      aws.String("RAW"),
		SigningAlgorithm: aws.String("ECDSA_SHA_256"),
	}
	result, err := svc.Sign(input)
	if err != nil {
		return nil, err
	}

	return result.Signature, nil
}

func (a *AwsKMSClient) Verify(keyID string, version string, message []byte, signature []byte) error {
	//svc := kms.New(a.Session)
	//input := &kms.VerifyInput{
	//	KeyId:            aws.String(keyID),
	//	Message:          message,
	//	MessageType:      aws.String("RAW"),
	//	Signature:        signature,
	//	SigningAlgorithm: aws.String("ECDSA_SHA_256"),
	//}
	//result, err := svc.Verify(input)
	//if err != nil {
	//	return err
	//}
	//if *result.SignatureValid {
	//	return nil
	//}

	key, err := a.GetKey(keyID, version)
	if err != nil {
		return err
	}

	pub, err := ParseECDSAPublicKey(key)
	if err != nil {
		return err
	}

	return VerifyECDSA(pub, message, signature)
}
