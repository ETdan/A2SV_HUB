package domain

// models/common.go

import "time"

type Country struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	ShortCode string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}

type Role struct {
	ID        int       `gorm:"primaryKey"`
	Type      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}

type Division struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
}

type File struct {
	ID        int       `gorm:"primaryKey"`
	URL       string    `gorm:"type:varchar(255)"`
	Name      string    `gorm:"type:varchar(255)"`
	Fid       string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}

type Fund struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	Amount      float32   `gorm:"type:real"`
	Currency    string    `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
}

type Event struct {
	ID        int       `gorm:"primaryKey"`
	Type      string    `gorm:"type:varchar(255)"`
	StartTime time.Time `gorm:"type:timestamp"`
	EndTime   time.Time `gorm:"type:timestamp"`
	Recurring int       `gorm:"type:integer"`
	Link      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}
