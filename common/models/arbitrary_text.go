package models

import (
	"encoding/json"
	"github.com/gynshu-one/goph-keeper/common/utils"
)

type ArbitraryText struct {
	// ArbitraryText is the text
	Text string `json:"text" bson:"text"`
}

func (data *ArbitraryText) EncryptAll(passphrase string) (encryptedData []byte, err error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return utils.EncryptData(marshaled, passphrase)

}

func (data *ArbitraryText) DecryptAll(passphrase string, encrypteData []byte) error {
	decrypted, err := utils.DecryptData(encrypteData, passphrase)
	if err != nil {
		return err
	}
	return json.Unmarshal(decrypted, data)
}