package dao

import "time"

type SignInInput struct {
	UserName string `binding:"required"`
	Password string `binding:"required"`
}

// Validation structs
type CreateStudentIntput struct {
	UserName  string     `binding:"required"`
	Password  string     `binding:"required"`
	FullName  string     `binding:"required"`
	Email     *string    `binding:"required"`
	Birthday  *time.Time `binding:"required"`
	ClassID   uint
	IsMonitor bool
}

type UpdateStudentInput struct {
	Password  string
	FullName  string
	Email     *string
	Birthday  *time.Time
	ClassID   uint
	IsMonitor bool
}
