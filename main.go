package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("CRUD.json")
	if err != nil {
		panic(err)
	}
	json := ParseJSON(data)
	tablesNames := json.TablesNames()
	fmt.Println(tablesNames)
	fmt.Println(json.FieldsOfTable("users").Create)
}
