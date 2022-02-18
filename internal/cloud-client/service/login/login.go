package login

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types/login"
	"github.com/golang-jwt/jwt"
)

var (
//signKey *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	tokenType string
	login.RequestLogin
}

func Login(request login.RequestLogin) (response login.ResponseLogin, err error) {
	response.Token, err = createToken(request)
	fatal(err)
	fmt.Println(response.Token)
	return
}

func createToken(request login.RequestLogin) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"level1",
		request,
	}
	signKey, err := rsa.GenerateKey(rand.Reader, 2048)
	fmt.Println(signKey.PublicKey)
	if err != nil {
		panic(err)
	}
	return t.SignedString(signKey)
}

func VerifyInvitation(request login.RequestInvitation) (response login.ResponseInvitation, err error) {

	return
}

func Sign(request login.RequestSign) (err error) {
	return
}
