package models

import (
	"gin_demo/entities"
	"time"
)

type Paging struct {
	Page     int `default:"0"`
	PageSize int `default:"10"`
}

type DeleteItemInput struct {
	IsDeleted bool      `default:"true"`
	DeletedAt time.Time `default:"time.Now()"`
}

type BaseModel interface {
	// Find
	GetById(id int) (entities.BaseEntity, error)
	GetByPage(page Paging) ([]entities.BaseEntity, error)

	// Create
	AddItem(baseEntity entities.BaseEntity) bool
	BulkAddItems(baseEntities []entities.BaseEntity) bool

	// Update
	UpdateItem(id int, baseEntity entities.BaseEntity) bool
	BulkUpdateItems(baseEntities []entities.BaseEntity) bool

	// Delete
	DeleteItem(id int) bool
	BulkDeleteItems(ids []int) bool
}
