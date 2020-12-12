package main
import (
  "books/Config"
  "books/Models"
  "books/Routes"
	"fmt"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	dsn := Config.DbUrl(Config.BuildDBConfig())
	Config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println("Status:", err)
	}

	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})
	 
	r := Routes.SetupRouter()
 	//running
 	r.Run()
}