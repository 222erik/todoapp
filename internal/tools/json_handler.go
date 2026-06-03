package tools

import (
	"bytes"
	"encoding/json"
	"io"
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

func RestoreTodo(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var ret []Task
	err = json.Unmarshal(jsonData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
