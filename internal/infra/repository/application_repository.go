package repository

import "gorm.io/gorm"

type DbApplicationRepository struct {
	DB *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *DbApplicationRepository {
	return &DbApplicationRepository{ DB: db }
}
