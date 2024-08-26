package local

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Todo struct {
	Title string
	Done  bool
}

var StorageFile = filepath.Join(os.Getenv("Desktop"), "todo.json")

func LoadTodos() ([]Todo, error) {
	var todos []Todo

	todos = append(todos, Todo{Title: "Buy Milk", Done: false})
	todos = append(todos, Todo{Title: "Buy Eggs", Done: true})
	todos = append(todos, Todo{Title: "Buy Bread", Done: false})

	//if the first is true, the second will run
	if _, err := os.Stat(StorageFile); os.IsNotExist(err) {
		return todos, nil
	}

	data, err := os.ReadFile(StorageFile) // returns a byte slice
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &todos)
	//todos becomes a slice of Todo

	if err != nil {
		return nil, err
	}

	return todos, nil
}
