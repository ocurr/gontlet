package gontlet

var (
	tableList map[string]*Table
	transport Transport
)

func InitServer(port string) {
	tableList = make(map[string]*Table)
	transport = newServer(port)
	go transport.serve()
	go updateTransport()
}

func InitClient(address string) {
	tableList = make(map[string]*Table)
	transport = newClient(address)
	go transport.serve()
	go updateTransport()
}

func GetTable(tableName string) *Table {
	if tableList[tableName] == nil {
		tableList[tableName] = &Table{Pairs: make([]*Entry, 0)}
	}
	return tableList[tableName]
}

func updateTransport() {
	for {
		for _, v := range tableList {
			for _, entry := range v.Pairs {
				if entry.Updated {
					transport.sendOutgoing([]byte(entry.Name + "=" + entry.Value))
					entry.Updated = false
				}
			}
		}
	}
}
