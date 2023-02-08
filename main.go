package main

import (
	_ "api/routers"
	"database/sql"
	"github.com/beego/beego/v2/client/orm"

	beego "github.com/beego/beego/v2/server/web"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func init() {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.DefaultTimeLoc = time.FixedZone("Asia/Ho_Chi_Minh", 7*60*60)

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, database)
	fmt.Println(dsn)
	orm.RegisterDataBase("default", "mysql", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
	beego.Run()
}
