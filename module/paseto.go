package module

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cerdas-buatan/be/model"
	"github.com/o1egl/paseto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Encode(id, role, secret string) (string, error) {
	now := time.Now()
	payload := model.Payload{
		ID:       id,
		Role:     role,
		IssuedAt: now,
		Expiry:   now.Add(24 * time.Hour),
	}

	v2 := paseto.NewV2()
	return v2.Encrypt([]byte(secret), payload, nil)
}

func Decode(secret, token string) (model.Payload, error) {
	var payload model.Payload
	v2 := paseto.NewV2()
	err := v2.Decrypt(token, []byte(secret), &payload, nil)
	if err != nil {
		return payload, err
	}

	if time.Now().After(payload.Expiry) {
		return payload, fmt.Errorf("token has expired")
	}

	return payload, nil
}

func Encode2(id primitive.ObjectID, role, privateKey string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.Set("id", id)
	token.SetString("role", role)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err
}

// func Decode2(publicKey string, tokenstring string) (payload model.Payload, err error) {
// 	var token *paseto.Token
// 	var pubKey paseto.V4AsymmetricPublicKey
// 	pubKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publicKey) // this wil fail if given key in an invalid format
// 	if err != nil {
// 		fmt.Println("Decode NewV4AsymmetricPublicKeyFromHex : ", err)
// 	}
// 	parser := paseto.NewParser()                                // only used because this example token has expired, use NewParser() (which checks expiry by default)
// 	token, err = parser.ParseV4Public(pubKey, tokenstring, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail
// 	if err != nil {
// 		fmt.Println("Decode ParseV4Public : ", err)
// 	} else {
// 		json.Unmarshal(token.ClaimsJSON(), &payload)
// 	}
// 	return payload, err
// }

func GenerateKey() (privateKey, publicKey string) {
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey = secretKey.Public().ExportHex()     // DO share this one
	privateKey = secretKey.ExportHex()
	return privateKey, publicKey
}

// package module

// import (
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"aidanwoods.dev/go-paseto"
// 	"github.com/cerdas-buatan/be/model"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func Encode(id primitive.ObjectID, role, privateKey string) (string, error) {
// 	token := paseto.NewToken()
// 	token.SetIssuedAt(time.Now())
// 	token.SetNotBefore(time.Now())
// 	token.SetExpiration(time.Now().Add(2 * time.Hour))
// 	token.Set("id", id)
// 	token.SetString("role", role)
// 	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
// 	return token.V4Sign(secretKey, nil), err
// }

func Decode(publicKey string, tokenstring string) (payload model.Payload, err error) {
	var token *paseto.Token
	var pubKey paseto.V4AsymmetricPublicKey
	pubKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publicKey) // this wil fail if given key in an invalid format
	if err != nil {
		fmt.Println("Decode NewV4AsymmetricPublicKeyFromHex : ", err)
	}
	parser := paseto.NewParser()                                  // only used because this example token has expired, use NewParser() (which checks expiry by default)
	token, err = parser.ParseV4Public(pubKey, tokenstring, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail
	if err != nil {
		fmt.Println("Decode ParseV4Public : ", err)
	} else {
		json.Unmarshal(token.ClaimsJSON(), &payload)
	}
	return payload, err
}

// func GenerateKey() (privateKey, publicKey string) {
// 	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
// 	publicKey = secretKey.Public().ExportHex()     // DO share this one
// 	privateKey = secretKey.ExportHex()
// 	return privateKey, publicKey
// }

// module/module.go

