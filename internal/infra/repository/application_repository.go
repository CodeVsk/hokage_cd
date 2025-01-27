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

func (r *DbApplicationRepository) GetBySlug(slug string) (*model.Application, error) {
	var application *model.Application
	result := r.DB.Where("slug = ?", slug).First(&application)
	if(result.Error != nil){
		return nil, result.Error
	}

	return application, nil
}

func (r *DbApplicationRepository) GetById(applicationId string) (*model.Application, error) {
	var application model.Application
	result := r.DB.Where("id = ?", applicationId).First(&application)
	if(result.Error != nil){
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil;
		}

		return nil, result.Error
	}

	return &application, nil
}

func (r *DbApplicationRepository) Create(model *model.Application) error {
	result := r.DB.Create(model)
	if(result.Error != nil){
		return result.Error
	}

	return nil
}

func (r *DbApplicationRepository) Delete(applicationId string) error {
	var application *model.Application
	result := r.DB.Where("id = ?", applicationId).Delete(&application)
	if(result.Error != nil){
		return result.Error
	}

	return nil
}

func (r *DbApplicationRepository) Update(model *model.Application) error {
	result := r.DB.Model(&model).Where("id = ?", model.ID).Updates(model)
	if(result.Error != nil){
		return result.Error
	}

	return nil
}