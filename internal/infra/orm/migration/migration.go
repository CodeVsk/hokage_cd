package migration

import (
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Application{}); err != nil {
		return err
	}

	return nil
}