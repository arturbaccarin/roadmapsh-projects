package task

import (
	"bytes"
	"encoding/json"
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

	tasks.Update(1, "Buy groceries and cook dinner")
	task = tasks[1]

	if task.Description != "Buy groceries and cook dinner" {
		t.Errorf("expected Description to be 'Buy groceries and cook dinner', got %s", task.Description)
	}

	if task.UpdatedAt == "" {
		t.Error("expected UpdatedAt, got empty")
	}

	tasks.Update(2, "Buy groceries and cook dinner")
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

	expected := "ID: 3\nDescription: Task 3\nStatus: open\n\nID: 1\nDescription: Task 1\nStatus: open\n\n"
	if buf.String() != expected {
		t.Errorf("expected: \n%s\ngot:\n%s", expected, buf.String())
	}
}

func TestString(t *testing.T) {
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

func TestChangeStatus(t *testing.T) {
	task := Task{
		ID:          1,
		Description: "Task 1",
		Status:      Open,
	}

	task.MarkAsDone()
	if task.Status != Done {
		t.Errorf("expected Done, got %s", task.Status)
	}

	task.MarkAsInProgress()
	if task.Status != InProgress {
		t.Errorf("expected InProgress, got %s", task.Status)
	}

	task.MarkAsOpen()
	if task.Status != Open {
		t.Errorf("expected Open, got %s", task.Status)
	}
}

func TestLoadFromJSONFile(t *testing.T) {
	tests := []struct {
		name      string
		file      string
		wantError bool
		wantTasks int
	}{
		{
			name:      "Valid JSON",
			file:      "valid_tasks.json", // This file should contain valid JSON
			wantError: false,
			wantTasks: 2, // Based on the number of tasks in valid_tasks.json
		},
		{
			name:      "File Not Found",
			file:      "nonexistent_file.json",
			wantError: true,
			wantTasks: 0,
		},
		{
			name:      "Invalid JSON",
			file:      "invalid_json.json", // This should contain malformed JSON
			wantError: true,
			wantTasks: 0,
		},
		{
			name:      "Empty File",
			file:      "empty_file.json", // Empty or empty array in JSON
			wantError: false,
			wantTasks: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a Tasks instance
			tasks := Tasks{}

			// Create a temporary file with the mock JSON data
			mockData := ""
			switch tt.file {
			case "valid_tasks.json":
				mockData = `[{"id": 1, "description": "Do laundry", "status": "pending", "created_at": "2025-05-11", "updated_at": "2025-05-11"}, {"id": 2, "description": "Buy groceries", "status": "done", "created_at": "2025-05-10", "updated_at": "2025-05-10"}]`
			case "invalid_json.json":
				mockData = `[{"id": 1, "description": "Do laundry", "status": "pending", "created_at": "2025-05-11", "updated_at": "2025-05-11",}` // Invalid trailing comma
			case "empty_file.json":
				mockData = `[]`
			case "nonexistent_file.json":
				err := tasks.LoadFromJSONFile(tt.file)
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			// Create a temporary file with the mock data
			tmpFile, err := os.CreateTemp("", "test_*.json")
			if err != nil {
				t.Fatal("Unable to create temporary file:", err)
			}
			defer os.Remove(tmpFile.Name()) // Clean up the temp file after test

			// Write mock data to the temp file
			if _, err := tmpFile.Write([]byte(mockData)); err != nil {
				t.Fatal("Failed to write mock data:", err)
			}
			tmpFile.Close()

			// Load tasks from the file
			err = tasks.LoadFromJSONFile(tmpFile.Name())

			// Check if we expect an error
			if (err != nil) != tt.wantError {
				t.Errorf("Expected error: %v, got: %v", tt.wantError, err != nil)
			}

			// Check if the number of tasks is correct
			if len(tasks) != tt.wantTasks {
				t.Errorf("Expected %d tasks, got %d", tt.wantTasks, len(tasks))
			}
		})
	}
}

func TestToJSON(t *testing.T) {
	tasks := Tasks{
		1: {
			ID:          1,
			Description: "Task 1",
			Status:      Open,
		},
		2: {
			ID:          2,
			Description: "Task 2",
			Status:      Done,
		},
	}

	ts := []Task{
		{
			ID:          1,
			Description: "Task 1",
			Status:      Open,
		},
		{
			ID:          2,
			Description: "Task 2",
			Status:      Done,
		},
	}

	expected, _ := json.Marshal(ts)
	actual, _ := tasks.ToJSON()

	if !bytes.Equal(expected, actual) {
		t.Errorf("Expected JSON: %s, got: %s", expected, actual)
	}
}
