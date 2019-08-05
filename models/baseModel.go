package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	Model
}

func (bm *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (bm *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
