package Models

import (
	"fmt"

	"../Config"
	_ "github.com/go-sql-driver/mysql" // To get ready function like create and save and where
)

// GetAllTodos return all todos
func GetAllTodos(todo *[]Todo) (err error) {
	if err := Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateATodo create a todo
func CreateATodo(todo *Todo) (err error) {
	if err := Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// GetATodo retrun a todo
func GetATodo(todo *Todo, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// UpdateATodo update a todo
func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

// DeleteATodo delete a todo
func DeleteATodo(todo *Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
