package model

import "time"

type EducationProgram struct {
	UpdatedAt              time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt              time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	EducationProgramID     uint      `gorm:"primaryKey;column:education_program_id" json:"id"`
	EducationProgramName   string    `gorm:"column:education_program_name" json:"name"`
	MaximalCountOfStudents *int      `gorm:"column:maximal_count_of_students;default:4" json:"maximal_count_of_students"`
}

func (o *EducationProgram) TableName() string {
	return "student_repo.education_program"
}
