package repository

import (
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
	"gorm.io/gorm"
)

type DbDeploymentRepository struct {
	DB *gorm.DB
}

func NewDeploymentRepository(db *gorm.DB) *DbDeploymentRepository {
	return &DbDeploymentRepository{ DB: db }
}

func (r *DbDeploymentRepository) GetById(deploymentId string) (*model.Deployment, error) {
	var deployment model.Deployment
	result := r.DB.Where("id = ?", deploymentId).First(&deployment)
	if(result.Error != nil){
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil;
		}

		return nil, result.Error
	}

	return &deployment, nil
}

func (r *DbDeploymentRepository) Create(model *model.Deployment) error {
	result := r.DB.Create(model)
	if(result.Error != nil){
		return result.Error
	}

	return nil
}

func (r *DbDeploymentRepository) Delete(deploymentId string) error {
	var deployment *model.Deployment
	result := r.DB.Where("id = ?", deploymentId).Delete(&deployment)
	if(result.Error != nil){
		return result.Error
	}

	return nil
}