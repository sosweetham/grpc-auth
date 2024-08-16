package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/sohamjaiswal/grpc-ftp/tools"
)

func main() {
	if err := tools.ValidateEnv(); err != nil {
		log.Warnf(": %v", err)
	}
	log.Print("Hi")
}