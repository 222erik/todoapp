package tools

import (
	"encoding/json"
	"os"
)

func SaveTodo(todo []any, filename string) {
	json_data, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(json_data)
	if err != nil {
		panic(err)
	}
}
