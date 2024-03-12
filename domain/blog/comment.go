package blog

import (
 "time"
 "gorm.io/gorm"
 "github.com/google/uuid"
)

type Comment struct {
	ID        string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID    string         `json:"user_id" gorm:"type:char(36)"`
	PostID    string         `json:"post_id" gorm:"type:char(36);not null"`
	Name      string         `json:"name" gorm:"type:char(50);not null"`
	Comment   string         `json:"comment" gorm:"type:text"`
	Likes     uint32         `json:"likes"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
