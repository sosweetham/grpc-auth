package tools

import (
	"errors"
	"time"

	"github.com/o1egl/paseto"
	"github.com/sohamjaiswal/grpc-ftp/pkg/payload"
)

type Tokenizer struct {
	Paseto 			*paseto.V2
	SymmetricKey	[]byte
}

func (tokenizer *Tokenizer) CreateAuthToken(username string, duration time.Duration) (*string, *payload.AuthPayload, error) {
	payload, err := payload.NewAuthPayload(username, duration)
	if err != nil {
		return nil, payload, err
	}
	token, err := tokenizer.Paseto.Encrypt(tokenizer.SymmetricKey, payload, nil)
	if err != nil {
		return nil, payload, err
	}
	return &token, payload, nil
}

func (tokenizer *Tokenizer) VerifyAuthToken(token string) (*payload.AuthPayload, error) {
	payload := &payload.AuthPayload{}

	if err := tokenizer.Paseto.Decrypt(token, tokenizer.SymmetricKey, payload, nil); err != nil {
		return nil, errors.New("invalid token")
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}