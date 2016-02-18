package gontlet

import (
	//"fmt"
	"strconv"
)

type Entry struct {
	Name string
	Value string
	Updated bool
}

type Table struct {
	Pairs []*Entry
}

func (t *Table) GetEntry(name string) *Entry {
	for _,e := range t.Pairs {
		if e.Name == name {
			return e
		}
	}
	return nil
}

func (t *Table) Update(name, value string) {
	e := t.GetEntry(name)
	if e != nil {
		e.Value = value
		e.Updated = true
	} else {
		t.Pairs = append(t.Pairs, &Entry{Name: name, Value: value, Updated: true})
	}
}

func (t *Table) getAsString(key string) (string, bool) {
	v := t.GetEntry(key)
	if v != nil {
		return v.Value, true
	}
	return "", false
}

func (t *Table) getAsInt(key string) (int64, bool) {
	v := t.GetEntry(key)
	if v != nil {
		i, err := strconv.ParseInt(v.Value,10,32)
		if err == nil {
			return i, true
		}
	}
	return 0, false
}

func (t *Table) getAsBool(key string) (bool,bool) {
	v := t.GetEntry(key)
	if v != nil {
		b, err := strconv.ParseBool(v.Value)
		if err == nil {
			return b, true
		}
	}
	return false,false
}

func (t *Table) getAsDouble(key string) (float64, bool) {
	v := t.GetEntry(key)
	if v != nil {
		d, err := strconv.ParseFloat(v.Value,64)
		if err == nil {
			return d, true
		}
	}
	return 0,false
}

