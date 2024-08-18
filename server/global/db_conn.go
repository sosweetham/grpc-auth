package global

import (
	"os"

	"github.com/sohamjaiswal/grpc-auth/pkg/storage"
	"gorm.io/gorm"
)

var defaultDB *gorm.DB

func setupConnection() (*gorm.DB, error) {
	connection := &storage.Db{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}
	conn, err := connection.NewConnection()
	if err != nil {
		return nil, err
	}
	defaultDB = conn
	return defaultDB, nil
}

func GetDBConn(restartConnection bool) (*gorm.DB,error) {
	if restartConnection {
		conn, err := setupConnection()
		if err != nil {
			return nil, err
		}
		defaultDB = conn
		return defaultDB, nil
	}
	if defaultDB == nil {
		conn, err := setupConnection()
		if err != nil {
			return nil, err
		}
		defaultDB = conn
		return defaultDB, nil
	}
	return defaultDB, nil
}
