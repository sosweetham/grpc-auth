package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Host		string
	Port 		string	
	Password	string
	User		string
	DBName 		string
	SSLMode 	string
}

func (db *Db) NewConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host,
		db.Port,
		db.User,
		db.Password,
		db.DBName,
		db.SSLMode,
	)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});
	if err != nil {
		return connection, err
	}
	return connection, nil
}