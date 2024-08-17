package payload

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AuthPayload struct {
	ID        uuid.UUID `json:"id"`
	Username  string
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expired_at"`
}

func NewAuthPayload(username string, duration time.Duration) (*AuthPayload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &AuthPayload{
		ID:        tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}, nil
}

func (payload *AuthPayload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return errors.New("token expired")
	}
	return nil
}