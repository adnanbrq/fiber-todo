package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoToString(t *testing.T) {
	todo := Todo{
		ID:          1,
		Description: "Test",
		Done:        false,
		CreatedAt:   0,
		UpdatedAt:   0,
	}

	expected := "Todo<id: 1, description: Test, done: false, created_at: 0>"
	assert.Equal(t, expected, todo.ToString())
}
