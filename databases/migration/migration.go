package main

import (
	"github.com/BoomTHDev/golang_clean_arch/config"
	"github.com/BoomTHDev/golang_clean_arch/databases"
	"github.com/BoomTHDev/golang_clean_arch/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	tx := db.ConnectionGetting().Begin()

	userMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return
	}
}

func userMigration(tx *gorm.DB) {
	tx.AutoMigrate(&entities.User{})
}
