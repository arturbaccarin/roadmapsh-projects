package task

import (
	"bytes"
	"io"
	"os"
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

func TestRemove(t *testing.T) {
	task := NewTask("Buy groceries")

	tasks := Tasks{}
	tasks.Add(&task)

	tasks.Remove(1)

	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tasks))
	}
}

func TestListAll(t *testing.T) {
	task1 := Task{
		ID:          1,
		Description: "Task 1",
		Status:      Open,
	}

	task2 := Task{
		ID:          2,
		Description: "Task 2",
		Status:      InProgress,
	}

	tasks := Tasks{
		1: task1,
		2: task2,
	}

	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	tasks.ListAll()

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	expected := "ID: 1\nDescription: Task 1\nStatus: open\n\nID: 2\nDescription: Task 2\nStatus: in_progress\n\n"
	if buf.String() != expected {
		t.Errorf("expected: \n%s\ngot:\n%s", expected, buf.String())
	}
}

func TestListByStatus(t *testing.T) {
	task1 := Task{
		ID:          1,
		Description: "Task 1",
		Status:      Open,
	}

	task2 := Task{
		ID:          2,
		Description: "Task 2",
		Status:      InProgress,
	}

	task3 := Task{
		ID:          3,
		Description: "Task 3",
		Status:      Open,
	}

	tasks := Tasks{
		1: task1,
		2: task2,
		3: task3,
	}

	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	tasks.ListByStatus(Open)

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	expected := "ID: 1\nDescription: Task 1\nStatus: open\n\nID: 3\nDescription: Task 3\nStatus: open\n\n"
	if buf.String() != expected {
		t.Errorf("expected: \n%s\ngot:\n%s", expected, buf.String())
	}
}

func TestStringI(t *testing.T) {
	task := Task{
		ID:          1,
		Description: "Task 1",
		Status:      Open,
	}

	expected := "ID: 1\nDescription: Task 1\nStatus: open\n"
	if task.String() != expected {
		t.Errorf("expected %s, got %s", expected, task.String())
	}
}
