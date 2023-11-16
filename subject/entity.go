package subject

import (
	"time"

	"github.com/muhammadputraage/goschool-api/student"
	"github.com/muhammadputraage/goschool-api/teacher"
)

type Entity struct {
	ID          string           `json:"id" gorm:"primaryKey"`
	SubjectName string           `json:"subject_name"`
	Teachers    []teacher.Entity `json:"teachers" gorm:"foreignKey:TeacherNumber;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,OnDelete:SET NULL;"`
	Students    []student.Entity `json:"students" gorm:"many2many:subject_students;foreignKey:ID;joinForeignKey:SubjectId;References:ID;joinReferences:StudentId;constraint:OnUpdate:CASCADE;"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

func (Entity) TableName() string {
	return "subjects"
}
