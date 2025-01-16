package repository

import (
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
	"gorm.io/gorm"
)

type DbApplicationRepository struct {
	DB *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *DbApplicationRepository {
	return &DbApplicationRepository{ DB: db }
}

func (r *DbApplicationRepository) Create(model *model.Application) error {
	result := r.DB.Create(model)
	if(result.Error != nil){
		return result.Error
	}

	return nil
}