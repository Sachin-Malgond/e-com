package main 

import(
	"fmt"
	"amz1-main/config"
	"amz1-main/model"
	"amz1-main/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	config.DB,err=gorm.Open("mysql",config.DbConnectionString(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Error Connecting to Infix DB: ",err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&model.Product{})

	r:=rest.SetupRouter()

	r.Run()
}