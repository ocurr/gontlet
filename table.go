package gontlet

import (
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

func (t *Table) GetEntry(name string) *Entry {
	for _,v := range t.Pairs {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (t *Table) Update(e Entry) {
	v := t.GetEntry(e.Name)
	if v != nil {
		v.Value = e.Value
	} else {
		t.Pairs = append(t.Pairs, e)
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

