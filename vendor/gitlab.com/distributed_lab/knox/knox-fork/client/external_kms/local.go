package external_kms

import (
	"fmt"
)

type LocalClient struct {
	Keys map[string]interface{}
}

func NewLocalClient(rawCfg map[string]interface{}) (Client, error) {
	client := LocalClient{}

	client.Keys = rawCfg["keys"].(map[string]interface{})

	//priv, err := ParseED25519PrivateKey([]byte(client.Keys["ed25519"].(string)))
	//if err != nil {
	//	panic(err)
	//}
	//
	//signature, err := SignED25519(*priv, []byte("bebra"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("signature:", string(signature))
	//
	//err = VerifyED25519(priv.Public().(ed25519.PublicKey), []byte("bebra"), signature)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Verified!")

	return &client, nil
}

func (l *LocalClient) GetKey(keyID string, version string) ([]byte, error) {
	key := l.Keys[keyID]

	if key == nil {
		return nil, fmt.Errorf("Not found")
	}

	return []byte(key.(string)), nil
}

func (l *LocalClient) Sign(keyID string, version string, message []byte) ([]byte, error) {
	key, err := l.GetKey(keyID, version)
	if err != nil {
		return nil, err
	}

	priv, err := ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	return Sign(priv, message)
}

func (l *LocalClient) Verify(keyID string, version string, message []byte, signature []byte) error {
	key, err := l.GetKey(keyID, version)
	if err != nil {
		return err
	}

	priv, err := ParsePrivateKey(key)
	if err != nil {
		return err
	}

	return Verify(priv, message, signature)
}
