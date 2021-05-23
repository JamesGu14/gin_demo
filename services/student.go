package services

import (
	"errors"
	"gin_demo/dao"
	"gin_demo/models"
	"strconv"
)

var studentModel models.StudentModel = models.StudentModel{}

func FindStudents(paging models.Paging) []dao.Student {
	return studentModel.GetStudents(paging)
}

func FindStudent(studentIdStr string) (dao.Student, error) {
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		return dao.Student{}, errors.New("invalid student id")
	}
	student, err := studentModel.GetStudent(studentId)
	if err != nil {
		return dao.Student{}, errors.New("student does not exist")
	}
	return student, nil
}

func CreateStudent(newStudent dao.Student) {
	studentModel.CreateStudent(newStudent)
}

func UpdateStudent(student dao.Student, input dao.UpdateStudentInput) {
	models.UpdateStudent(student, input)
}

func DeleteStudent(student_id int) {
	models.DeleteStudent(student_id)
}
