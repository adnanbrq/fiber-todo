package dto

import "github.com/adnanbrq/fiber-todo/model"

// InputCreateTodo specifies data used to create a todo
type InputCreateTodo struct {
	Description interface{} `valid:"required|string|min:1"`
	Done        interface{} `valid:"required|bool"`
}

// ToMap returns DTO's data as a map
func (i InputCreateTodo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Description": i.Description.(string),
		"Done":        i.Done.(bool),
	}
}

// ToTodo will create a new todo object containing the DTO Data
func (i InputCreateTodo) ToTodo() model.Todo {
	return model.Todo{
		Description: i.Description.(string),
		Done:        i.Done.(bool),
	}
}

// InputUpdateTodo specifies data used to update a todo
type InputUpdateTodo struct {
	Description interface{} `valid:"nullable|string|min:1"`
	Done        interface{} `valid:"nullable|bool"`
}

// ToMap returns DTO's data as a map
func (i InputUpdateTodo) ToMap() map[string]interface{} {
	changes := make(map[string]interface{})

	if i.Description != nil {
		changes["Description"] = i.Description.(string)
	}

	if i.Done != nil {
		changes["Done"] = i.Done.(bool)
	}

	return changes
}
