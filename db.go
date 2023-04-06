package main

import (
	"database/sql"
	"fmt"

	"github.com/go-mysql-org/go-mysql/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initMysql() mysql.Position {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v", DbUser, DbPassWord, DbAddress, DbName))
	if err != nil {
		panic(err)
	}

	// Get least binlog pos
	var pos mysql.Position
	var binlogDoDB, binlogIgnoreDB, executedGTIDSet string
	if err = db.QueryRow("SHOW MASTER STATUS").Scan(&pos.Name, &pos.Pos, &binlogDoDB, &binlogIgnoreDB, &executedGTIDSet); err != nil {
		panic(err)
	}

	DB = db
	return pos
}
