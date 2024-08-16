package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/sohamjaiswal/grpc-ftp/api"
	"github.com/sohamjaiswal/grpc-ftp/global"
	"github.com/sohamjaiswal/grpc-ftp/models"
	"github.com/sohamjaiswal/grpc-ftp/tools"
)


func initializeDb() *gorm.DB {
	db, err := global.GetDBConn(false)
	for err != nil {
		log.Error("didnt get connection again...")
		time.Sleep(1 * time.Second)
		db, err = global.GetDBConn(false)
	}
	log.Info("DB connection success!")
	return db
}

func main() {
	if err := tools.ValidateEnv(); err != nil {
		log.Warnf(": %v", err)
	}
	db := initializeDb()
	if err := models.MigrateUser(db); err != nil {
		log.Fatal("could not migrate user DB!")
	}
	api.Start()
}