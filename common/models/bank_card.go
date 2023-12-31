package models

import (
	"encoding/json"
	"github.com/gynshu-one/goph-keeper/common/utils"
)

// BankCard is a struct for bank card
type BankCard struct {
	// Info is the additional info about the card
	Info string `json:"info" bson:"info"`
	// CardType is the type of card such as Visa, MasterCard, etc.
	CardType string `json:"card_type" bson:"cardType"`
	// CardNum is the card number
	CardNum string `json:"card_num" bson:"cardNum"`
	// CardName is the name on the card
	CardName string `json:"card_name" bson:"cardName"`
	// CardCvv is the card's CVV
	CardCvv string `json:"card_cvv" bson:"cardCvv"`
	// CardExp is the card's expiration date
	CardExp string `json:"card_exp" bson:"cardExp"`
}

// EncryptAll encrypts all sensitive fields
func (data *BankCard) EncryptAll(passphrase string) (encryptedData []byte, err error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return utils.EncryptData(marshaled, passphrase)
}

// DecryptAll decrypts all sensitive fields
func (data *BankCard) DecryptAll(passphrase string, encrypteData []byte) error {
	decrypted, err := utils.DecryptData(encrypteData, passphrase)
	if err != nil {
		return err
	}
	return json.Unmarshal(decrypted, &data)
}
