package main

import (
	"github.com/karyono7528/go-auth-jwt/app"
)

func main() {
	// db, _ = gorm.Open("mysql", "luxsoft:k0piiT2mku@tcp(127.0.0.1:3306)/db_golang?charset=utf8&parseTime=True&loc=Local")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()
	// db.AutoMigrate(&Person{})
	app.StartApplication()
}
