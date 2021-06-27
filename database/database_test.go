package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

type User struct {
	ID   int
	Name string
}

// 配合数据驱动包-github.com/go-sql-driver/mysql
// jinzhu （gorm）底层也是 dabase/sql + mysql驱动包的封装
// dabase/sql/driver 包定义了应被数据库驱动实现的接口，这些接口会被 sql 包使用
// github.com/go-sql-driver/mysql 包是对 "database/sql/driver"实现
var dbPool *sql.DB

func TestDataBase(t *testing.T) {
	dbPool, err := NewPoolDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	rows, err := dbPool.Query("select * from user where id = ?", 1)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close() // 注意这里，一定要关闭
	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatal(err)
			return
		}
		break
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func NewPoolDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "xxxxxx"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "mydb"
	dsn := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
