package gontlet

import (
	//"fmt"
)

var (
	tableList map[string]*Table
	server *Server
)

func Init() {
	tableList = make(map[string]*Table)
	server = newServer("8081")
	go server.serve()
	go updateServer()
}

func GetTable(tableName string) *Table {
	if tableList[tableName] == nil {
		tableList[tableName] = &Table{Pairs: make([]*Entry, 0)}
	}
	return tableList[tableName]
}

func updateServer() {
	for {
		for _,v := range tableList {
			for _,entry := range v.Pairs {
				if entry.Updated {
					server.sendOutgoing([]byte(entry.Name+"="+entry.Value))
					entry.Updated = false
				}
			}
		}
	}
}
