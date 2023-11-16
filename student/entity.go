package student

import (
	"time"
)

type Entity struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	StudentName   string    `json:"student_name"`
	StudentNumber string    `json:"student_number" gorm:"unique"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Entity) TableName() string {
	return "students"
}
