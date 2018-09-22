package main

import (
	"encoding/json"
)

// General Comment
type General struct {
	Host   string            `json:"database_host"`
	User   string            `json:"database_user"`
	Pass   string            `json:"database_pass"`
	Name   string            `json:"database_name"`
	Type   string            `json:"database_type"`
	Tables map[string]tables `json:"tables"`
}

type tables struct {
	Create []string `json:"create"`
	Read   []string `json:"read"`
	Update []string `json:"update"`
	Delete []string `json:"delete"`
}

// ParseJSON : Returns a General Object, from a json object
func ParseJSON(data []byte) General {
	var jsonObject General
	json.Unmarshal(data, &jsonObject)
	return jsonObject
}

// TablesNames : Get tables names from a General type
func (gen *General) TablesNames() []string {
	tables := make([]string, len(gen.Tables))
	i := 0
	for key := range gen.Tables {
		tables[i] = key
		i++
	}
	return tables
}

// FieldsOfTable : Returns the fields of CRUD for a table in a General struct
func (gen *General) FieldsOfTable(table string) tables {
	return gen.Tables[table]
}
