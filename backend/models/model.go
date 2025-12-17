package models

import "time"

type Link struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	OriginalURL string `json:"original_url"`
	ShortCode string `gorm:"unique;not null;index" json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
}

func (l *Link) TableName() string{
	return "link"
}