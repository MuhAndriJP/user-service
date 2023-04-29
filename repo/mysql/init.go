package mysql

import (
	"fmt"
	"os"

	"github.com/MuhAndriJP/user-service.git/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	usernameAndPassword := fmt.Sprint(os.Getenv("db_user")) + ":" + fmt.Sprint(os.Getenv("db_password"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("db_database")) + "?charset=utf8&parseTime=true&loc=UTC&tls=true"
	// config := os.Getenv("CONNECTION_DB")

	var e error
	DB, e = gorm.Open(mysql.Open(urlConnection), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	InitMigrate()
}

func InitMigrate() {
	// DB.Migrator().DropTable(&entity.Users{})
	DB.AutoMigrate(&entity.Users{})
}
