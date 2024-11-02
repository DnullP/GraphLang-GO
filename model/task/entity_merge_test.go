package task_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/DnullP/GraphLang-GO/model/task"
)

// Mocking the GlobelModel.Input function
type MockModel struct{}

func (m *MockModel) Input(prompt string) string {
	// Mock response based on the input prompt
	if strings.Contains(prompt, "Alice Bob") {
		return `{
			"Alice": ["Alice", "Alicia"],
			"Bob": ["Bob", "Bobby"]
		}`
	}
	return "{}"
}

func TestMergeEntities(t *testing.T) {
	// Replace the real model with the mock model

	tests := []struct {
		name     string
		entities []string
		expected map[string]interface{}
	}{
		{
			name:     "Test with Alice and Bob",
			entities: []string{"Alice", "Alicia", "Bob", "Bobby"},
			expected: map[string]interface{}{
				"Alice": []interface{}{"Alice", "Alicia"},
				"Bob":   []interface{}{"Bob", "Bobby"},
			},
		},
		{
			name:     "Test with empty list",
			entities: []string{"岩永", "岩永小姐", "岩永琴子", "九郎先生", "樱川九郎", "琴子", "纱季", "九郎"},
			expected: map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := task.MergeEntities(tt.entities)
			fmt.Println(result)
		})
	}
}
