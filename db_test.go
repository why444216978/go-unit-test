package main

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_handle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	// (.+) 用于替代字段，可用于 select、order、group等
	mock.ExpectQuery("SELECT (.+) from test where id > ?").WillReturnRows(sqlmock.NewRows([]string{"id", "goods_id", "name"}).AddRow(1, 1, "1"))
	mock.ExpectExec("UPDATE test SET goods_id").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO test").WithArgs(1, "1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = handle(db); err != nil {
		panic(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		panic(err)
	}
}

func Test_handleOrm(t *testing.T) {
	db := NewMemoryDB()
	err := db.Migrator().CreateTable(&Test{})
	assert.Nil(t, err)

	handleOrm(db)
}

func NewMemoryDB() *gorm.DB {
	var db *gorm.DB
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	dialector := sqlite.Open(":memory:?cache=shared")
	if db, err = gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		panic(err)
	}
	dba, err := db.DB()
	dba.SetMaxOpenConns(1)
	return db
}

func CloseMemoryDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
