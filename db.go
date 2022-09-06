package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/test"

type Test struct {
	ID      int64  `json:"id" db:"id" gorm:"column:id"`
	GoodsID int64  `json:"goodsID" db:"goods_id" gorm:"column:goods_id"`
	Name    string `json:"name" db:"name" gorm:"column:name"`
}

func (Test) TableName() string {
	return "test"
}

func handle(db *sql.DB) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * from test where id > ?", 0)
	if err != nil {
		panic(err)
	}
	result := []Test{}
	if err = sqlx.StructScan(rows, &result); err != nil {
		panic(err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println("sql:", string(b))

	if _, err = tx.Exec("UPDATE test SET goods_id = goods_id + 1 where id = 2"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO test (goods_id, name) VALUES (?, ?)", 1, "1"); err != nil {
		return
	}
	return
}

func main() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = handle(db); err != nil {
		panic(err)
	}

	orm, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	handleOrm(orm)
}

func handleOrm(orm *gorm.DB) {
	var rows []Test

	clause := func(db *gorm.DB) *gorm.DB {
		return db.Where("id >= ?", 1)
	}
	err := clause(orm.Select("*")).Find(&rows).Error
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(rows)
	if err != nil {
		panic(err)
	}
	fmt.Println("gorm", string(b))
}
