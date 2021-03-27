package models

import (
	"gin_demo/entities"
	"gin_demo/util"
	"github.com/fatih/structs"
)

type StudentModel struct{}

func (s *StudentModel) GetById(id int) (entities.BaseEntity, error) {
	var student entities.Student
	if err := util.DB.Where("id = ?", id).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (s *StudentModel) GetByPage(page Paging) ([]entities.BaseEntity, error) {
	var students []entities.Student
	util.DB.Limit(page.PageSize).Offset(page.PageSize * page.PageSize).Find(&students)
	baseEntities := make([]entities.BaseEntity, 0)
	for _, e := range students {
		baseEntities = append(baseEntities, e)
	}
	return baseEntities, nil
}

func (s *StudentModel) AddItem(baseEntity entities.BaseEntity) bool {
	newStudent := baseEntity.(entities.Student)
	util.DB.Create(&newStudent)
	return true
}

func (s *StudentModel) BulkAddItems(baseEntities []entities.BaseEntity) bool {
	panic("implement me")
}

func (s *StudentModel) UpdateItem(id int, baseEntity entities.BaseEntity) bool {
	_update := structs.Map(baseEntity)
	util.DB.Model(&entities.Student{}).Where("id = ?", id).Updates(_update)
	return true
}

func (s *StudentModel) BulkUpdateItems(baseEntities []entities.BaseEntity) bool {
	panic("implement me")
}

func (s *StudentModel) DeleteItem(id int) bool {
	util.DB.Delete(&entities.Student{}, id)
	return true
}

func (s *StudentModel) BulkDeleteItems(ids []int) bool {
	panic("implement me")
}
