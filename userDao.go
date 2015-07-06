package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insertUser(name string, pwd string) (int, bool) {
	log.Println("Open Mysql Connection")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sample?charset=utf8")
	if err != nil {
		log.Fatal("Fail to connect : ", err)
		return 0, false
	}

	defer db.Close()

	log.Println("prepare statment")
	stmt, err := db.Prepare("INSERT INTO user(name, pwd) VALUES( ?, ? )")

	if err != nil {
		log.Fatal("prepare statment error : ", err)
		return 0, false
	}
	defer stmt.Close()

	log.Println("excute statment")
	res, err := stmt.Exec(name, pwd)
	if err != nil {
		log.Fatal(err)
		return 0, false
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, false
	}

	return int(lastId), true
}

func selectUser(name string, pwd string) bool {
	log.Println("Open Mysql Connection")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sample?charset=utf8")
	if err != nil {
		log.Fatal("Fail to connect : ", err)
		return false
	}

	defer db.Close()

	log.Println("prepare statment")
	stmtOut, err := db.Prepare("SELECT * FROM user WHERE name = ? and pwd =? ")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return false
	}
	defer stmtOut.Close()

	user := &User{}

	err = stmtOut.QueryRow(name, pwd).Scan(&user)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return false
	}

	log.Println("%s, %s", user.name, user.pwd)

	if user.name != "" && user.pwd != "" {
		return true
	} else {
		return false
	}
}
