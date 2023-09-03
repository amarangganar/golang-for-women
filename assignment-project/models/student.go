package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name" gorm:"not null;type:varchar(300)" valid:"required~Name is required"`
	Age       int     `json:"age" gorm:"not null" valid:"required~Age is required,type(int)~Age must be an integer"`
	Scores    []Score `json:"scores" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (student *Student) BeforeUpdate(tx *gorm.DB) (err error) {
	// since we have no way of knowing which score item is updated
	// the approach is to remove the student's existing scores first
	// then added the newly updated scores
	// this assuming that the student's old scores are obsolete, hence the update
	tx.Where("student_id = ?", student.ID).Delete(&Score{})

	return
}
