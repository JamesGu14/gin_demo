package models

import (
	"gin_demo/entities"
	"gin_demo/util"
	"github.com/fatih/structs"
)

type StudentModel struct{}

func (sm *StudentModel) GetStudents(paging Paging) []entities.Student {
	var students []entities.Student
	util.DB.Limit(paging.PageSize).Offset(paging.PageSize * paging.PageSize).Find(&students)
	return students
}

func GetStudent(id int) (entities.Student, error) {
	var student entities.Student
	if err := util.DB.Where("id = ?", id).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func CreateStudent(newStudent entities.Student) {
	util.DB.Create(&newStudent)
}

type UpdateStudentInput2 struct {
	Password string
}

func UpdateStudent(existingStudent entities.Student, updateStudent entities.UpdateStudentInput) {
	//DB.Model(&existingStudent).Updates(updateStudent)
	//_birthDay := time.Date(1990, 5, 11, 20, 34, 58, 0, time.UTC)
	_update := structs.Map(updateStudent)
	util.DB.Model(&existingStudent).Updates(_update)
}

func DeleteStudent(student_id int) {
	util.DB.Delete(&entities.Student{}, student_id)
}

func BulkDeleteStudent(studentIds []int) {

}
