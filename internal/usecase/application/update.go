package usecase

import (
	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
	"github.com/codevsk/codevsk_golang_cd/pkg/result"
	"github.com/google/uuid"
)

type UpdateApplicationUseCase struct {
	Uow uow.UnitOfWork
}

func NewUpdateApplicationUseCase(uow uow.UnitOfWork) *UpdateApplicationUseCase {
	return &UpdateApplicationUseCase{
		Uow: uow,
	}
}

func (c *UpdateApplicationUseCase) Execute(applicationId string,  input dto.UpdateApplicationInputDTO) (result.Result) {
	if(applicationId == ""){
		return result.NewValidationResult("The application id cannot be null")
	}

	if _, err := uuid.Parse(applicationId); err != nil {
		return result.NewValidationResult("The application ID is not valid")
	}

	application, err := c.Uow.GetApplicationRepository().GetById(applicationId)
	if err != nil {
		return result.NewBadRequestResult("Application not found for the id entered")
	}

	applicationEntity := model.ToApplicationEntity(application)

	if(input.Name != "") {
		applicationEntity.SetName(input.Name)
	}

	if(input.Slug != "") {
		applicationEntity.SetSlug(input.Slug)
	}

	if(input.Repository_url != "") {
		applicationEntity.SetRepositoryUrl(input.Repository_url)
	}

	if err := applicationEntity.Validate(); err != nil {
		return result.NewValidationResult(err.Error())
	}

	applicationEntity.SetUpdated("master")

	application = model.ToApplicationModel(applicationEntity)

	if err := c.Uow.GetApplicationRepository().Update(application); err != nil {
		if err := c.Uow.Rollback(); err != nil {
			return result.NewInternalResult(err)
		}

		return result.NewInternalResult(err)
	}

	if err := c.Uow.Commit(); err != nil {
		return result.NewInternalResult(err)
	}

	return result.NewNoContentResult()
}