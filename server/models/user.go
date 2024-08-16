package models

import (
	"time"

	"github.com/sohamjaiswal/grpc-ftp/global"
	"github.com/sohamjaiswal/grpc-ftp/tools"
	"gorm.io/gorm"
)

type User struct {
	ID string `gorm:"primary key; type:uuid; default:gen_random_uuid(); uniqueIndex" json:"id"`
	Username *string `gorm:"type:varchar(40); not null; uniqueIndex" json:"username"`
	Password *string `gorm:"size:255; not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MigrateUser(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hash, err := global.DefaultHasher.GenerateFromPassword(*u.Password)
	if err != nil {
		return err
	}
	u.Password = &hash
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	existingUser := User{}
	err = tx.Where("id = ?", u.ID).First(&existingUser).Error
	if err != nil {
		return err
	}
	match, err := tools.ComparePasswordAndHash(*u.Password, *existingUser.Password)
	if err != nil {
		return err
	}
	if !match {
		hash, err := global.DefaultHasher.GenerateFromPassword(*u.Password)
		if err != nil {
			return err
		}
		u.Password = &hash
	}
	return nil
}
