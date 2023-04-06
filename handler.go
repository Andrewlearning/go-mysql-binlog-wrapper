package main

import (
	"encoding/json"
	"fmt"
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
	rowData := make(map[string]interface{})
	for i, col := range e.Table.Columns {
		rowData[col.Name] = e.Rows[0][i]
	}

	jsonData, err := json.Marshal(rowData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	return nil
}
