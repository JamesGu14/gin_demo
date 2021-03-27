package models

import "time"

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
	GetById(id int)
	GetByPage(page Paging)

	// Create
	AddItem()
	BulkAddItems()

	// Update
	UpdateItem()
	BulkUpdateItems()

	// Delete
	DeleteItem()
	BulkDeleteItems()
}
