package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var db *sql.DB

func main() {
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/emp?charset=utf8")
	db, err := sql.Open("mysql", "root:root@/emp?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()

	CheckErr(err)
	rows, err := db.Query("select * from loginlist")
	CheckErr(err)
	for rows.Next() {
		var id int
		var LoginID string
		var LoginNa string
		var Password string
		err = rows.Scan(&id, &LoginID, &LoginNa, &Password)
		CheckErr(err)
		fmt.Println(id, LoginID, LoginNa, Password)
	}
	db.Close()
	go signalListen()
	for {
		time.Sleep(10 * time.Second)
		fmt.Println("running... pid:", os.Getpid())
	}

}
func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) //让进程收集信号量。
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
	}
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
