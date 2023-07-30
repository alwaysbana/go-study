package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	DBname := "yuanshen"
	timoout := "10s"
	//  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timoout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败,err%" + err.Error())
	}
	fmt.Println(db)
}
func main() {

}
