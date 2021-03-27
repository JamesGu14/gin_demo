package entities

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity interface{}

type Teacher struct {
	gorm.Model
	UserName  string `gorm:"unique_index"`
	Password  string
	FullName  string
	Email     *string
	Birthday  *time.Time
	TeachYear int8
	Courses   []Course
	IsDeleted bool
}

type Student struct {
	gorm.Model
	UserName  string `gorm:"unique_index"`
	Password  string
	FullName  string
	Email     *string
	Birthday  *time.Time
	ClassID   uint
	IsMonitor bool
	Courses   []*Course `gorm:"many2many:student_courses;"`
	IsDeleted bool
}

type Class struct {
	gorm.Model
	ClassName string
	Grade     int8
	TeacherID uint `gorm:"column:coordinator"`
	Students  []Student
}

type Course struct {
	gorm.Model
	CourseName string
	TeacherID  uint
}
