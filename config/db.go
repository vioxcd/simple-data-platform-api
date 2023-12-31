package config

import (
	"errors"
	"vioxcd/dpl/models"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("fail to load file")
	}
}

func ConnectToDB() {
	var dbConfig DBConfig = DBConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	fmt.Println(dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Database Connection Error")
	}

	migration()
}

func migration() {
	DB.AutoMigrate(&models.UserLog{})

	if err := DB.AutoMigrate(&models.Run{}); err == nil && DB.Migrator().HasTable(&models.Run{}) {
		if err := DB.First(&models.Run{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&models.Run{
				Type:        "All",
				Description: "Run Snapshot for Data",
			})
		}
	}

	if err := DB.AutoMigrate(&models.Role{}); err == nil && DB.Migrator().HasTable(&models.Role{}) {
		if err := DB.First(&models.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			roles := []models.Role{
				{Id: 1, Name: "Admin"},
				{Id: 2, Name: "User"},
			}
			DB.Create(&roles)
		}
	}

	if err := DB.AutoMigrate(&models.User{}); err == nil && DB.Migrator().HasTable(&models.User{}) {
		if err := DB.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&models.User{
				Id:   1,
				Name: "Admin",
				RoleId: 1,
				Email: "admin@app.com",
				Password: "admin",
			})
		}
	}
}
