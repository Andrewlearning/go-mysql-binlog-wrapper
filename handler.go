package main

import (
	"log"

	"github.com/go-mysql-org/go-mysql/canal"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	rowDataOld := make(map[string]interface{})
	rowDataNew := make(map[string]interface{})
	for i, col := range e.Table.Columns {
		rowDataOld[col.Name] = e.Rows[0][i]
		rowDataNew[col.Name] = e.Rows[1][i]
	}

	log.Printf("[binlog][old] %v", rowDataOld)
	log.Printf("[binlog][new] %v", rowDataNew)

	return nil
}
