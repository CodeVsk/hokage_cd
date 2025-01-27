package usecase

import (
	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	"github.com/codevsk/codevsk_golang_cd/internal/entity"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
	"github.com/codevsk/codevsk_golang_cd/pkg/result"
)

type CreateApplicationUseCase struct {
	Uow uow.UnitOfWork
}

func NewCreateApplicationUseCase(uow uow.UnitOfWork) *CreateApplicationUseCase {
	return &CreateApplicationUseCase{
		Uow: uow,
	}
}

func (c *CreateApplicationUseCase) Execute(input dto.CreateApplicationInputDTO) (result.Result) {
	application , err := entity.NewApplication(input.Name, input.Slug, input.Repository_url, "master")
	if(err != nil){
		return result.NewValidationResult(err.Error())
	}

	if _, err := c.Uow.GetApplicationRepository().GetBySlug(application.Slug); err == nil {
		return result.NewConflictResult()
	}

	model := model.ToApplicationModel(application)

	if err := c.Uow.GetApplicationRepository().Create(model); err != nil {
		c.Uow.Rollback()

		return result.NewInternalResult(err)
	}

	if err := c.Uow.Commit(); err != nil {
		return result.NewInternalResult(err)
	}

	return result.NewCreatedResult()
}