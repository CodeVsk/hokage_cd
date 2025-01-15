package orm

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewGormSqlServerDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open("Server=localhost;Database=hokage;User Id=sa;Password=dev1234@;Trusted_Connection=True;"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	return db, nil
}