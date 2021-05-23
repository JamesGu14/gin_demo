package models

import (
	"gin_demo/dao"
	"gin_demo/util"
	"github.com/fatih/structs"
)

type StudentModel struct{}

func (sm *StudentModel) GetStudents(paging Paging) []dao.Student {
	var students []dao.Student
	util.DB.Limit(paging.PageSize).Offset(paging.PageSize * paging.PageSize).Find(&students)
	return students
}

func (sm *StudentModel) GetStudent(id int) (dao.Student, error) {
	var student dao.Student
	if err := util.DB.Where("id = ?", id).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (sm *StudentModel) CreateStudent(newStudent dao.Student) {
	util.DB.Create(&newStudent)
}

type UpdateStudentInput2 struct {
	Password string
}

func UpdateStudent(existingStudent dao.Student, updateStudent dao.UpdateStudentInput) {
	//DB.Model(&existingStudent).Updates(updateStudent)
	//_birthDay := time.Date(1990, 5, 11, 20, 34, 58, 0, time.UTC)
	_update := structs.Map(updateStudent)
	util.DB.Model(&existingStudent).Updates(_update)
}

func DeleteStudent(student_id int) {
	util.DB.Delete(&dao.Student{}, student_id)
}

func BulkDeleteStudent(studentIds []int) {

}
