package models

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	GroupName    string `gorm:"type:varchar(255)" json:"group_name,omitempty"`
	GroupPicture string `gorm:"type:text" json:"group_picture,omitempty"`
	Users        []User `gorm:"many2many:user_conversations;" json:"users,omitempty"`
}
