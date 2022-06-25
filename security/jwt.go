package security

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v4"
	"log"
	"os"
)

type Claim struct {
	jwt.RegisteredClaims
	UserId int64  `json:"user_id"`
	Auth   string `json:"auth"`
}

func Encode(id int64, Auth string) (*string, error) {
	key := []byte(os.Getenv("TOKEN_SECRET"))
	signer, err1 := jwt.NewSignerHS(jwt.HS256, key)
	if err1 != nil {
		log.Printf("signer err \nerr: %v", err1)
		return nil, err1
	}
	claim := &Claim{
		RegisteredClaims: jwt.RegisteredClaims{Subject: fmt.Sprintf("%d", id)},
		UserId:           id,
		Auth:             Auth,
	}
	builder := jwt.NewBuilder(signer)
	token, err2 := builder.Build(claim)
	if err2 != nil {
		log.Printf("jwt builder err \nerr: %v", err2)
		return nil, err2
	}
	var result = token.String()

	return &result, nil
}

func EncodeUser(id int64) (*string, error) {
	return Encode(id, "user")
}

func Decode(token string) (*Claim, error) {
	key := []byte(os.Getenv("TOKEN_SECRET"))
	verifier, err1 := jwt.NewVerifierHS(jwt.HS256, key)
	if err1 != nil {
		log.Printf("verifier err \nerr: %v", err1)
		return nil, err1
	}

	verifierToken, err2 := jwt.Parse([]byte(token), verifier)
	if err2 != nil {
		log.Printf("jwt, verifier err \nerr: %v", err2)
		return nil, err2
	}

	if err3 := verifier.Verify(verifierToken); err3 != nil {
		log.Printf("verify err \nerr: %v", err2)
		return nil, err3
	}

	var claim Claim
	if err3 := json.Unmarshal(verifierToken.Claims(), &claim); err3 != nil {
		log.Printf("verify err \nerr: %v", err3)
		return nil, err3
	}

	return &claim, nil
}

func GetUserId(ctx context.Context) int64 {
	claim := ctx.Value("claim").(*Claim)
	return claim.UserId
}

func GetClaim(ctx context.Context) *Claim {
	return ctx.Value("claim").(*Claim)
}
