package models

import (
	"gin_demo/entities"
	"gin_demo/common"
	"github.com/fatih/structs"
)

type StudentModel struct{}

func (s *StudentModel) GetById(id int) (entities.BaseEntity, error) {
	var student entities.Student
	if err := common.DB.Where("id = ?", id).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (s *StudentModel) GetByPage(page Paging) ([]entities.BaseEntity, error) {
	var students []entities.Student
	common.DB.Limit(page.PageSize).Offset(page.PageSize * page.PageSize).Find(&students)
	baseEntities := make([]entities.BaseEntity, 0)
	for _, e := range students {
		baseEntities = append(baseEntities, e)
	}
	return baseEntities, nil
}

func (s *StudentModel) AddItem(baseEntity entities.BaseEntity) bool {
	newStudent := baseEntity.(entities.Student)
	common.DB.Create(&newStudent)
	return true
}

func (s *StudentModel) BulkAddItems(baseEntities []entities.BaseEntity) bool {
	panic("implement me")
}

func (s *StudentModel) UpdateItem(id uint, baseEntity entities.BaseEntity) bool {
	_update := structs.Map(baseEntity)
	common.DB.Model(&entities.Student{}).Where("id = ?", id).Updates(_update)
	return true
}

func (s *StudentModel) BulkUpdateItems(baseEntities []entities.BaseEntity) bool {
	panic("implement me")
}

func (s *StudentModel) DeleteItem(id uint) bool {
	common.DB.Delete(&entities.Student{}, id)
	return true
}

func (s *StudentModel) BulkDeleteItems(ids []uint) bool {
	panic("implement me")
}
