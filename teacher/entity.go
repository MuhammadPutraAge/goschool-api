package teacher

import (
	"time"

	"github.com/muhammadputraage/goschool-api/student"
)

type Entity struct {
	ID            string           `json:"id" gorm:"primaryKey"`
	TeacherName   string           `json:"teacher_name"`
	TeacherNumber string           `json:"teacher_number" gorm:"unique"`
	SubjectID     string           `json:"subject_id"`
	Students      []student.Entity `json:"students" gorm:"many2many:teacher_students;foreignKey:ID;joinForeignKey:TeacherId;References:ID;joinReferences:StudentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}

func (Entity) TableName() string {
	return "teachers"
}
