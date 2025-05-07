package task

import (
	"testing"
)

func TestAdd(t *testing.T) {
	task := NewTask("Buy groceries")

	tasks := Tasks{}
	tasks.Add(&task)

	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}

	if task.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", task.ID)
	}

	if task.CreatedAt == "" {
		t.Error("expected CreatedAt, got empty")
	}

	if task.UpdatedAt != "" {
		t.Errorf("expected UpdatedAt to be empty, got %s", task.UpdatedAt)
	}
}

func TestUpdate(t *testing.T) {
	task := NewTask("Buy groceries")

	tasks := Tasks{}
	tasks.Add(&task)

	err := tasks.Update(1, "Buy groceries and cook dinner")
	task = tasks[1]

	if task.Description != "Buy groceries and cook dinner" {
		t.Errorf("expected Description to be 'Buy groceries and cook dinner', got %s", task.Description)
	}

	if task.UpdatedAt == "" {
		t.Error("expected UpdatedAt, got empty")
	}

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	err = tasks.Update(2, "Buy groceries and cook dinner")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
