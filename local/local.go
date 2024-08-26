package local

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Todo struct {
	Title   string
	Done    bool
	Created string
}

func checkIndex(index int, todos []Todo) error {
	if index > len(todos) {
		return errors.New("Index out of range")
	}
	return nil
}

var StorageFile = filepath.Join(os.TempDir(), "todos.json")

func LoadTodos() ([]Todo, error) {
	var todos []Todo

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

func SaveTodos(todo Todo) error {
	//json insert todo
	todos, err := LoadTodos()
	if err != nil {
		return err
	}
	todos = append(todos, todo)

	data, err := json.MarshalIndent(todos, "", "  ")
	//indent with 2 spaces

	if err != nil {
		return err
	}

	return os.WriteFile(StorageFile, data, 0644)
}

func RemoveTodos(index int) error {
	todos, err := LoadTodos()
	if err != nil {
		return err
	}

	index--

	err = checkIndex(index, todos)
	if err != nil {
		return err
	}

	for i := range todos {
		if i == index {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(StorageFile, data, 0644)
}

func UpdateTodos(index int, done bool) error {
	todos, err := LoadTodos()
	if err != nil {
		return err
	}

	index--
	err = checkIndex(index, todos)
	if err != nil {
		return err
	}

	todos[index].Done = done

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(StorageFile, data, 0644)
}
