package usecase

import (
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
	"github.com/codevsk/codevsk_golang_cd/pkg/result"
	"github.com/google/uuid"
)

type DeleteApplicationUseCase struct {
	Uow uow.UnitOfWork
}

func NewDeleteApplicationUseCase(uow uow.UnitOfWork) *DeleteApplicationUseCase {
	return &DeleteApplicationUseCase{
		Uow: uow,
	}
}

func (c *DeleteApplicationUseCase) Execute(applicationId string) (result.Result) {
	if(applicationId == ""){
		return result.NewValidationResult("The application id cannot be null")
	}

	if _, err := uuid.Parse(applicationId); err != nil {
		return result.NewValidationResult("The application ID is not valid")
	}

	if _, err := c.Uow.GetApplicationRepository().GetById(applicationId); err != nil {
		return result.NewBadRequestResult("Application not found for the id entered")
	}

	if err := c.Uow.GetApplicationRepository().Delete(applicationId); err != nil {
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