package services

import (
	"errors"
	"gin_demo/entities"
	"gin_demo/models"
	"strconv"
)

func FindStudents(paging models.Paging) []entities.Student {
	return models.GetStudents(paging)
}

func FindStudent(studentIdStr string) (entities.Student, error) {
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		return entities.Student{}, errors.New("invalid student id")
	}
	student, err := models.GetStudent(studentId)
	if err != nil {
		return entities.Student{}, errors.New("student does not exist")
	}
	return student, nil
}

func CreateStudent(newStudent entities.Student) {
	models.CreateStudent(newStudent)
}

func UpdateStudent(student entities.Student, input entities.UpdateStudentInput) {
	models.UpdateStudent(student, input)
}

func DeleteStudent(student_id int) {
	models.DeleteStudent(student_id)
}
