package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	bob := `{"name":"bob","age":33}`
	var b Person

	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)

	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}
