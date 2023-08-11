package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string `gorm:"type:varchar(100);unique_index" json:"username"`
	Email         string `gorm:"type:varchar(100);unique_index" json:"email"`
	PhoneNumber   string `gorm:"type:varchar(100);unique_index" json:"phone_number,omitempty"`
	ProfilePicture string `gorm:"type:text" json:"profile_picture,omitempty"`
	HashedPassword []byte `gorm:"type:bytea" json:"-"`
	Conversations []Conversation `gorm:"many2many:user_conversations;" json:"conversations,omitempty"`
}
