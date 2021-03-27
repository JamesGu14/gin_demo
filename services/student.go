package services

import (
	"errors"
	"gin_demo/entities"
	"gin_demo/models"
	"strconv"
)

func FindStudents(paging models.Paging) ([]entities.Student, error) {
	students := make([]entities.Student, 0)
	_entities, err := new(models.StudentModel).GetByPage(paging)
	if err != nil {
		return nil, err
	}
	for _, e := range _entities {
		students = append(students, e.(entities.Student))
	}
	return students, nil
}

func FindStudent(studentIdStr string) (entities.Student, error) {
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		return entities.Student{}, errors.New("invalid student id")
	}
	student, err := new(models.StudentModel).GetById(studentId)
	if err != nil {
		return entities.Student{}, errors.New("student does not exist")
	}
	return student.(entities.Student), nil
}

func CreateStudent(newStudent entities.Student) {
	new(models.StudentModel).AddItem(newStudent)
}

//
//func UpdateStudent(student entities.Student, input entities.UpdateStudentInput) {
//	models.UpdateStudent(student, input)
//}
//
//func DeleteStudent(student_id int) {
//	models.DeleteStudent(student_id)
//}
