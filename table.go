package gontlet

import (
	"fmt"
	"strconv"
)

type Entry struct {
	Name string
	Value string
}

type Table struct {
	Name string
	Pairs []Entry
}

var tableList map[string][]*Table

func AddTable(name string) {
	tableList[name] = &Table{Name: name, Pairs: make(map[string][]byte)})
}

func UpdateTable(name string, e Entry) {
	tableList[name].Update(e);
}

func GetTable(name string) *Table {
	return tableList[name]
}

func (t *Table) GetEntry(name string) *Entry {
	for _,v in range t.Pairs {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (t *Table) Update(e Entry) {
	if v := t.GetEntry(e.Name) != nil {
		v.Value = e.Value
	} else {
		t.Pairs = append(t.Pairs, e)
	}
}

func (t *Table) getAsString(key string) string {
	if v := t.GetEntry(e.Name) != nil {
		return v.Value
	} else {
		return nil
	}
}

func (t *Table) getAsInt(key string) int32 {
	if v := t.GetEntry(e.Name) != nil {
		return strconv.ParseInt(v.Value,10,32)
	} else {
		return nil
	}
}

func (t *Table) getAsBool(key string) bool {
	if v := t.GetEntry(e.Name) != nil {
		return strconv.ParseBool(v.Value)
	} else {
		return nil
	}
}

func (t *Table) getAsDouble(key string) float64 {
	if v := t.GetEntry(e.Name) != nil {
		return strconv.ParseFloat(v.Value,64)
	} else {
		return nil
	}
}

