package main

import (
	"log"

	"github.com/go-mysql-org/go-mysql/canal"
)

func main() {
	c, _ := initCanal()

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
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
