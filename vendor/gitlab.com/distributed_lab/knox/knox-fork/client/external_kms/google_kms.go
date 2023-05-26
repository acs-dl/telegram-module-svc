package external_kms

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"gitlab.com/distributed_lab/figure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"hash/crc32"
	"sync"
)

type GoogleKMSClientConfig struct {
	Project  string `fig:"project"`
	Location string `fig:"location"`
	Keyring  string `fig:"keyring"`
}

type GoogleKMSClient struct {
	mutex  *sync.RWMutex
	Config *GoogleKMSClientConfig
	keys   map[string]map[string][]byte
}

func NewGoogleKMSClient(rawCfg map[string]interface{}) (Client, error) {
	config := GoogleKMSClientConfig{}
	err := figure.Out(&config).From(rawCfg).Please()
	if err != nil {
		return nil, err
	}
	keys := make(map[string]map[string][]byte)
	mutex := &sync.RWMutex{}
	return &GoogleKMSClient{
		mutex:  mutex,
		Config: &config,
		keys:   keys,
	}, nil
}

func (g *GoogleKMSClient) getCacheKey(keyID string, version string) []byte {
	g.mutex.RLock()
	cachedKey := g.keys[keyID][version]
	g.mutex.RUnlock()
	return cachedKey
}

func (g *GoogleKMSClient) setCacheKey(keyID string, version string, key []byte) {
	g.mutex.Lock()
	g.keys[keyID] = make(map[string][]byte)
	g.keys[keyID][version] = key
	g.mutex.Unlock()
}

func (g *GoogleKMSClient) getPath(keyID string, version string) string {
	return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", g.Config.Project, g.Config.Location, g.Config.Keyring, keyID, version)
}

func (g *GoogleKMSClient) GetKey(keyID string, version string) ([]byte, error) {
	cacheKey := g.getCacheKey(keyID, version)
	if cacheKey != nil {
		return cacheKey, nil
	}

	path := g.getPath(keyID, version)

	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	req := &kmspb.GetPublicKeyRequest{
		Name: path,
	}

	result, err := client.GetPublicKey(ctx, req)
	if err != nil {
		return nil, err
	}

	key := []byte(result.Pem)

	crc32c := func(data []byte) uint32 {
		t := crc32.MakeTable(crc32.Castagnoli)
		return crc32.Checksum(data, t)
	}
	if int64(crc32c(key)) != result.PemCrc32C.Value {
		return nil, err
	}

	g.setCacheKey(keyID, version, key)

	return key, nil
}

func (g *GoogleKMSClient) Sign(keyID string, version string, message []byte) ([]byte, error) {
	path := g.getPath(keyID, version)

	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create kms client: %v", err)
	}
	defer client.Close()

	digest := sha256.New()
	if _, err := digest.Write(message); err != nil {
		return nil, fmt.Errorf("failed to create digest: %v", err)
	}

	crc32c := func(data []byte) uint32 {
		t := crc32.MakeTable(crc32.Castagnoli)
		return crc32.Checksum(data, t)

	}
	digestCRC32C := crc32c(digest.Sum(nil))

	req := &kmspb.AsymmetricSignRequest{
		Name: path,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: digest.Sum(nil),
			},
		},
		DigestCrc32C: wrapperspb.Int64(int64(digestCRC32C)),
	}

	result, err := client.AsymmetricSign(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to sign digest: %v", err)
	}

	if result.VerifiedDigestCrc32C == false {
		return nil, fmt.Errorf("AsymmetricSign: request corrupted in-transit")
	}

	if int64(crc32c(result.Signature)) != result.SignatureCrc32C.Value {
		return nil, fmt.Errorf("AsymmetricSign: response corrupted in-transit")
	}

	return result.Signature, nil
}

func (g *GoogleKMSClient) Verify(keyID string, version string, message []byte, signature []byte) error {
	pubKey, err := g.GetKey(keyID, version)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(pubKey)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %v", err)
	}
	ecKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("public key is not elliptic curve")
	}

	digest := sha256.Sum256(message)
	ok = ecdsa.VerifyASN1(ecKey, digest[:], signature)
	if !ok {
		return fmt.Errorf("failed to verify signature")
	}

	return nil
}
