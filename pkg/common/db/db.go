package db

import (
	"fmt"
	"log"
	"simple-orm/pkg/common/config"
	"simple-orm/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	fmt.Println(c.DBUser)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})
	return db

}
