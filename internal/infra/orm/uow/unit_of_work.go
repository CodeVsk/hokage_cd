package uow

import (
	"github.com/codevsk/codevsk_golang_cd/internal/contract"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/repository"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	Commit() error
	Rollback() error
	GetApplicationRepository() contract.ApplicationRepository
}

type Uow struct {
	DB *gorm.DB
	Tx *gorm.DB
	ApplicationRepository contract.ApplicationRepository
}

func NewUnitOfWork(db *gorm.DB) (*Uow, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error // Caso a transação não possa ser iniciada, retorna o erro
	}

	return &Uow{
			DB: db,
			Tx: tx,
			ApplicationRepository: repository.NewApplicationRepository(tx),
		}, nil
}

func (o *Uow) Commit() error {
	if err := o.Tx.Commit().Error; err != nil {
		return err
	}

	return nil;
}

func (o *Uow) Rollback() error {
	if err := o.Tx.Rollback().Error; err != nil {
		return err
	}

	return nil
}

func (o *Uow) GetApplicationRepository() contract.ApplicationRepository {
	return o.ApplicationRepository
}