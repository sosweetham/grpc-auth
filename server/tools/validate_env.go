package tools

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func ValidateEnv() error {
	missing := []string{}
	envVars := []string{
		"HOST",
		"PORT", 
		"MODE",
		"DB_HOST",
		"DB_PORT",
		"DB_PASS",
		"DB_USER",
		"DB_NAME",
		"DB_SSLMODE",
		"PASS_SALT",
	}
	for _, currVar := range envVars {
		setVar := os.Getenv(currVar)
		log.Printf("%s = %s", currVar, setVar)
		if setVar == "" {
			missing = append(missing, currVar)
		}
	}
	if cap(missing) >= 1 {
		err := fmt.Sprintf("total_missing: %v, missing: %v", len(missing), missing)
		return errors.New(err)
	}
	return nil
}