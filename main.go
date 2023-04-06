package main

import (
	"log"

	"github.com/go-mysql-org/go-mysql/canal"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initLogger()
	pos := initMysql()
	c, _ := initCanal()

	c.SetEventHandler(&MyEventHandler{})
	c.RunFrom(pos)

	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}

	defer release()
}

func initCanal() (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = DbAddress
	cfg.User = DbUser
	cfg.Password = DbPassWord
	cfg.Dump.TableDB = DbName
	cfg.Dump.Tables = []string{TableName}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return c, nil
}

func release() {
	logFile.Close()
	DB.Close()
}
