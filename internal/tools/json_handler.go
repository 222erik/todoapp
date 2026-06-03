package tools

import (
	"bytes"
	"encoding/json"
	"os"
)

type Task struct {
	TaskMessage string `json:"task"`
	Priority    uint8  `json:"priority"`
}

// SaveTodo saves the tasks in a Json file
func SaveTodo(todos []Task, filename string) error {
	jsonData, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	{
		var jsonDataBuffer bytes.Buffer
		json.Indent(&jsonDataBuffer, jsonData, "", "  ")
		jsonData = jsonDataBuffer.Bytes()
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
