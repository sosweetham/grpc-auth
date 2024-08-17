package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID 				uuid.UUID 	`gorm:"primary key; type:uuid; uniqueIndex" json:"id"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
	RefreshToken 	*string		`gorm:"type:varchar; uniqueIndex; not null" json:"refresh_token"`
	UserAgent 		*string 	`gorm:"type:varchar; not null" json:"user_agent"`
	ClientIp 		*string 	`gorm:"type:varchar; not null" json:"client_ip"`
	IsBlocked 		bool 		`gorm:"type:boolean; default false; not null" json:"is_blocked"`
	ExpiresAt 		time.Time 	`gorm:"type:timestamptz; not null" json:"expires_at"`
	Username 		*string 	`json:"username"`
}

func MigrateSession(db *gorm.DB) error {
	if err := db.AutoMigrate(&Session{}); err != nil {
		return err
	}
	return nil
}

func (s *Session) BeforeSave(tx *gorm.DB) (error) {
	existingSession := Session{}
	err := tx.Where("ID = ?", s.ID).First(&existingSession).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}