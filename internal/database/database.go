package database

import (
	"fmt"
	"gameReview/internal/config"
	"gameReview/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Review{}, &model.Review{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
