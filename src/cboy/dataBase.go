package cboy

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	url string
	numbers int
}

func GetAllData(db *sql.DB)  {
	rows, e := db.Query("select * from test")
	if e == nil {
		errors.New("query incur error")
		return
	}
	for rows.Next(){
		var user User
		e := rows.Scan(&user.url,&user.numbers)
		if e!=nil {
			break
		}
		fmt.Println(user.url," ",user.numbers)
	}
	rows.Close()
}
func QueryAge(db *sql.DB,url string) int{
	var user User
	db.QueryRow("select * from user where url=?",url).Scan(&user.url,&user.numbers)
	return user.numbers
}
func MysqlStart()  *sql.DB{
	db, err := sql.Open("mysql", "root:Aa1072303326#@tcp(localhost:3306)/cboy") //声明数据库类型为mysql 使用用户名,密码,连接形式,链接地址和端口,指定数据库
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	log.Println("数据库连接成功")
	return db
}
func InsertData(db *sql.DB,url string){
	_, err := db.Exec("insert into user(url,numbers)values(?,1) on duplicate key update numbers=numbers+1;",url)
	if err != nil {
		log.Printf("update err")
	}
}
func getAge(url string)  int{
	db:=MysqlStart()
	defer db.Close()
	CreatTable(db)
	InsertData(db, url)
	return QueryAge(db,url)
}
func CreatTable(db *sql.DB)  {
	_, err :=db.Exec("create table if not exists user(url varchar(20) primary key not null,numbers int);")
	if err != nil {
		log.Printf("update err")
	}
}