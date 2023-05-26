package external_kms

import (
	"encoding/json"
	vault "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/knox/knox-fork"
	"gitlab.com/distributed_lab/knox/knox-fork/log"
	"strconv"
	"sync"
)

type VaultClient struct {
	mutex  *sync.RWMutex
	Client *vault.Client
	Config *VaultClientConfig
	keys   map[string]map[string][]byte
}

type VaultClientConfig struct {
	Endpoint string `fig:"endpoint"`
	Token    string `fig:"token"`
	Path     string `fig:"path"`
}

func NewVaultClient(rawCfg map[string]interface{}) (Client, error) {
	config := VaultClientConfig{}
	err := figure.Out(&config).From(rawCfg).Please()
	if err != nil {
		return nil, err
	}

	vaultCfg := vault.DefaultConfig()

	vaultCfg.Address = config.Endpoint
	client, err := vault.NewClient(vaultCfg)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
		return nil, err
	}

	client.SetToken(config.Token)

	keys := make(map[string]map[string][]byte)
	mutex := &sync.RWMutex{}
	return &VaultClient{
		mutex:  mutex,
		Client: client,
		Config: &config,
		keys:   keys,
	}, nil
}

func (v *VaultClient) getCacheKey(keyID string, version string) []byte {
	v.mutex.RLock()
	cachedKey := v.keys[keyID][version]
	v.mutex.RUnlock()
	return cachedKey
}

func (v *VaultClient) setCacheKey(keyID string, version string, key []byte) {
	v.mutex.Lock()
	v.keys[keyID] = make(map[string][]byte)
	v.keys[keyID][version] = key
	v.mutex.Unlock()
}

type vaultResponse struct {
	Data vaultResponseData `json:"data"`
}

type vaultResponseData struct {
	Keys []knox.Key `json:"keys"`
}

func (v *VaultClient) GetKey(keyID string, versionID string) ([]byte, error) {
	cacheKey := v.getCacheKey(keyID, versionID)
	if cacheKey != nil {
		return cacheKey, nil
	}

	response, err := v.Client.Logical().Read(v.Config.Path)
	if err != nil {
		return nil, err
	}
	if response != nil {
		vaultResp := vaultResponse{}
		rawResponseData, err := json.Marshal(&response.Data)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal vault response")
		}
		err = json.Unmarshal(rawResponseData, &vaultResp)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to unmarshal vault response")
		}

		for _, currentKey := range vaultResp.Data.Keys {
			if currentKey.ID == keyID {
				for _, currentKeyVersion := range currentKey.VersionList.GetActive() {
					if strconv.Itoa(int(currentKeyVersion.ID)) == versionID {
						keyBytes := currentKeyVersion.Data
						v.setCacheKey(keyID, versionID, keyBytes)
						return keyBytes, nil
					}
				}
			}
		}
	}

	return nil, errors.New("Not found key")
}

func (v *VaultClient) Sign(keyID string, version string, message []byte) ([]byte, error) {
	key, err := v.GetKey(keyID, version)
	if err != nil {
		return nil, err
	}

	priv, err := ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	return Sign(priv, message)
}

func (v *VaultClient) Verify(keyID string, version string, message []byte, signature []byte) error {
	key, err := v.GetKey(keyID, version)
	if err != nil {
		return err
	}

	priv, err := ParsePrivateKey(key)
	if err != nil {
		return err
	}

	return Verify(priv, message, signature)
}
