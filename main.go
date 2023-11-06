package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	personMap := map[string]interface{}{
		"name":    "John",
		"age":     30,
		"married": true,
		"house": map[string]interface{}{
			"address": "21 2nd Street",
			"city":    "New York",
			"state":   "NY",
			"zip":     10021,
		},
	}

	jsonData, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json data %s\n", string(jsonData))
}
