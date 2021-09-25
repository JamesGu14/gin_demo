package services

import (
	"errors"
	"gin_demo/entities"
	"gin_demo/models"
	"gin_demo/common"
	"strconv"
	"context"
)

var ctx = context.Background()

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
	// Redis Cache operation
	err := common.R.Set(ctx, "find_student", 1, 0).Err()
	if err != nil {
		panic(err)
	}

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

func UpdateStudent(id uint, input entities.UpdateStudentInput) {
	new(models.StudentModel).UpdateItem(id, input)
}

func DeleteStudent(studentId uint) {
	new(models.StudentModel).DeleteItem(studentId)
}
