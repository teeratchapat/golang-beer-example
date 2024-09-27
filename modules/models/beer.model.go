package models

import "time"

type Beer struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	Type      string    `gorm:"type:varchar(255)"`
	Details   string    `gorm:"type:text"`
	ImageURL  string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty"`
	BeerID    uint      `bson:"beer_id"`
	Action    string    `bson:"action"`
	Timestamp time.Time `bson:"timestamp"`
	Data      Beer      `bson:"data"`
}
