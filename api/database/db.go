package database

import (
	"fmt"
	"os"

	"github.com/KaiqueSantosDev/gocrud/api/core/price"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	env := godotenv.Load()
	if env != nil {
		fmt.Print(env)
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")

	dbconnect := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", user, pass, dbhost, dbname)
	fmt.Println("Conectado")

	conn, err := gorm.Open(mysql.Open(dbconnect))
	if err != nil {
		fmt.Println(err)
	}
	DB = conn
	DB.AutoMigrate(&price.Product{})
}
