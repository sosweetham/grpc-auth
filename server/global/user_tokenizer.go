package global

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"github.com/sohamjaiswal/grpc-auth/tools"
)

func checkUserTokenDurationsValidity() (time.Duration, time.Duration) {
	atStrDur := os.Getenv("USER_ACCESS_TOKEN_DURATION")
	atDur, err := time.ParseDuration(atStrDur)
	if err != nil {
		log.Fatal("unparseable user access token validity duration")
	}
	rfStrDur := os.Getenv("USER_REFRESH_TOKEN_DURATION")
	rfDur, err:= time.ParseDuration(rfStrDur)
	if err != nil {
		log.Fatal("unparseable user refresh token validity duration")
	}
	return atDur, rfDur
}

func getUserAuthTokenizer() (*tools.Tokenizer) {
	tokenizerKey := os.Getenv("USER_TOKEN_SECRET")
	if len(tokenizerKey) != chacha20poly1305.KeySize {
		log.Fatalf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	return &tools.Tokenizer{
		Paseto: paseto.NewV2(),
		SymmetricKey: []byte(tokenizerKey),
	}
}

var GlobalUserAccessTokenDuration, 
	GlobalUserRefreshTokenDuration = checkUserTokenDurationsValidity()
var GlobalUserAuthTokenizer  = getUserAuthTokenizer()

