package models

import "time"

type Score struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	AssignmentTitle string `json:"assignment_title" gorm:"not null; type:varchar(300)" valid:"required~Assignment Title is required"`
	Description     string `json:"description" gorm:"not null; type:text"`
	Score           int    `json:"score" gorm:"not null" valid:"required~Score is required,type(int)~Score must be integer"`
	StudentID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
