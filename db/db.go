package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseSQL() *sql.DB {
	paramofdb := "root:Yjdsqgfhjkm10@tcp(127.0.0.1:3306)/magazin"

	db, err := sql.Open("mysql", paramofdb)

	if err != nil {
		panic(err)
	}
	//fmt.Println("успешно всё")
	//defer db.Close()

	return db
}

func DatabaseGorm() *gorm.DB {
	paramofdb := "root:Yjdsqgfhjkm10@tcp(127.0.0.1:3306)/magazin"

	db, err := gorm.Open(mysql.Open(paramofdb), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//fmt.Println("успешно всё")
	//defer db.Close()

	return db
}
