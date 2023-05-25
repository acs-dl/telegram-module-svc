package external_kms

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/knox/knox-fork/log"
	"math/big"
	"sync"
)

type AzureVaultClient struct {
	mutex  *sync.RWMutex
	Client *azkeys.Client
	Config *AzureVaultClientConfig
	keys   map[string]map[string][]byte
}

type AzureVaultClientConfig struct {
	VaultURI string `fig:"vault_uri"`
}

func NewAzureVaultClient(rawCfg map[string]interface{}) (Client, error) {
	config := AzureVaultClientConfig{}
	err := figure.Out(&config).From(rawCfg).Please()
	if err != nil {
		return nil, err
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return nil, err
	}

	client, err := azkeys.NewClient(config.VaultURI, cred, nil)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]map[string][]byte)
	mutex := &sync.RWMutex{}
	return &AzureVaultClient{
		mutex:  mutex,
		Client: client,
		Config: &config,
		keys:   keys,
	}, nil
}

func (a *AzureVaultClient) getCacheKey(keyID string, version string) []byte {
	a.mutex.RLock()
	cachedKey := a.keys[keyID][version]
	a.mutex.RUnlock()
	return cachedKey
}

func (a *AzureVaultClient) setCacheKey(keyID string, version string, key []byte) {
	a.mutex.Lock()
	a.keys[keyID] = make(map[string][]byte)
	a.keys[keyID][version] = key
	a.mutex.Unlock()
}

func (a *AzureVaultClient) GetKey(keyID string, version string) ([]byte, error) {
	cacheKey := a.getCacheKey(keyID, version)
	if cacheKey != nil {
		return cacheKey, nil
	}

	resp, err := a.Client.GetKey(context.Background(), keyID, version, nil)
	if err != nil {
		return nil, err
	}

	json, err := resp.MarshalJSON()
	fmt.Print(string(json))
	key, err := parseJSONWebKey(resp.Key)
	a.setCacheKey(keyID, version, key)
	return key, err
}

func parseJSONWebKey(key *azkeys.JSONWebKey) ([]byte, error) {
	x := big.NewInt(0).SetBytes(key.X)
	y := big.NewInt(0).SetBytes(key.Y)
	fmt.Println("x:", x.String(), "y:", y.String())

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(&pub)
	if err != nil {
		return nil, err
	}
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubASN1})

	return pemEncodedPub, nil
}

func (a *AzureVaultClient) Sign(keyID string, version string, message []byte) ([]byte, error) {
	algorithm := azkeys.JSONWebKeySignatureAlgorithmES256
	digest := sha256.New()
	if _, err := digest.Write(message); err != nil {
		return nil, fmt.Errorf("Failed to create digest: %v", err)
	}
	msgHashSum := digest.Sum(nil)

	params := azkeys.SignParameters{
		Algorithm: &algorithm,
		Value:     msgHashSum,
	}

	sigResponse, err := a.Client.Sign(context.Background(), keyID, version, params, nil)
	if err != nil {
		return nil, err
	}
	res, err := sigResponse.MarshalJSON()
	fmt.Println(string(res))

	return sigResponse.Result, nil
}

func (a *AzureVaultClient) Verify(keyID string, version string, message []byte, signature []byte) error {
	algorithm := azkeys.JSONWebKeySignatureAlgorithmES256
	digest := sha256.New()
	if _, err := digest.Write(message); err != nil {
		return fmt.Errorf("Failed to create digest: %v", err)
	}
	msgHashSum := digest.Sum(nil)

	params := azkeys.VerifyParameters{
		Algorithm: &algorithm,
		Digest:    msgHashSum,
		Signature: signature,
	}

	verifyResponse, err := a.Client.Verify(context.Background(), keyID, version, params, nil)
	if err != nil {
		return err
	}

	if !(*verifyResponse.Value) {
		return fmt.Errorf("failed to verify signature")
	}
	return nil

}
