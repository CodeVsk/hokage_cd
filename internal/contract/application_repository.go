package contract

import (
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
)

type ApplicationRepository interface {
	Create(model *model.Application) error
	Update(model *model.Application) error
	Delete(applicationId string) error
	GetById(id string) (*model.Application, error)
	GetBySlug(slug string) (*model.Application, error)
}
