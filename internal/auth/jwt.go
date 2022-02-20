package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	typesClient "github.com/byteso/Xcloud/api/cloud-client/v1/types"
	typesServer "github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/golang-jwt/jwt"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

type CustomClaimsClient struct {
	*jwt.StandardClaims
	TokenType string
	typesClient.RequestLogin
}

type CustomClaimsServer struct {
	*jwt.StandardClaims
	TokenType string
	typesServer.RequestLogin
}

func GenerateKeyForRsa() error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	signKey = key
	verifyKey = &key.PublicKey

	return nil
}

func CreateToken(content interface{}, types string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	switch types {
	case "client":
		t.Claims = &CustomClaimsClient{
			&jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			},
			"level1",
			func() typesClient.RequestLogin {
				if v, ok := content.(typesClient.RequestLogin); ok {
					return v
				}
				return typesClient.RequestLogin{}
			}(),
		}
	case "server":
		t.Claims = &CustomClaimsServer{
			&jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			},
			"level1",
			func() typesServer.RequestLogin {
				if v, ok := content.(typesServer.RequestLogin); ok {
					return v
				}
				return typesServer.RequestLogin{}
			}(),
		}
	}

	return t.SignedString(signKey)
}

func ParseToken(tokenString string, types string) bool {
	switch types {
	case "client":
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsClient{}, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if err != nil {
			log.Fatal(err)
		}

		if claims, ok := token.Claims.(*CustomClaimsClient); ok && token.Valid {
			fmt.Println(claims)
			return true
		}
	case "server":
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsServer{}, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(token)

		if claims, ok := token.Claims.(*CustomClaimsServer); ok && token.Valid {
			fmt.Println(claims.LoginCode)
			return true
		}
	}

	return false
}
