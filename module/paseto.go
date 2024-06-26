package module

import (
	"encoding/json"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/serbaevents/backendSE/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Encode(id primitive.ObjectID, role, privateKey string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.Set("id", id)
	token.SetString("role", role)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err
}

func Decode(publicKey string, tokenstring string) (payload model.Payload, err error) {
	var token *paseto.Token
	var pubKey paseto.V4AsymmetricPublicKey
	pubKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publicKey) 
	if err != nil {
		fmt.Println("Decode NewV4AsymmetricPublicKeyFromHex : ", err)
	}
	parser := paseto.NewParser()                            
	token, err = parser.ParseV4Public(pubKey, tokenstring, nil) 
	if err != nil {
		fmt.Println("Decode ParseV4Public : ", err)
	} else {
		json.Unmarshal(token.ClaimsJSON(), &payload)
	}
	return payload, err
}

func GenerateKey() (privateKey, publicKey string) {
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey = secretKey.Public().ExportHex()     // DO share this one
	privateKey = secretKey.ExportHex()
	return privateKey, publicKey
}