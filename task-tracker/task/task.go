package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type status string

const (
	Open       status = "open"
	InProgress status = "in_progress"
	Done       status = "done"
)

type Tasks map[int]Task

func (t Tasks) Add(task *Task) {
	task.ID = len(t) + 1
	task.CreatedAt = time.Now().String()

	t[task.ID] = *task
}

func (t Tasks) Update(id int, description string) error {
	task, ok := t[id]
	if !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}

	task.Description = description
	task.UpdatedAt = time.Now().String()

	t[id] = task

	return nil
}

func (t Tasks) Remove(id int) {
	delete(t, id)
}

func (t Tasks) ListAll() {
	for _, task := range t {
		fmt.Printf("%s\n", task)
	}
}

func (t Tasks) ListByStatus(status status) {
	for _, task := range t {
		if task.Status == status {
			fmt.Printf("%s\n", task)
		}
	}
}

func (t Tasks) LoadFromJSONFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	var tasks []Task
	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	for _, task := range tasks {
		t.Add(&task)
	}

	return nil
}

type Task struct {
	ID          int
	Description string
	Status      status
	CreatedAt   string
	UpdatedAt   string
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %d\nDescription: %s\nStatus: %s\n", t.ID, t.Description, t.Status)
}

func NewTask(description string) Task {
	return Task{
		Description: description,
		Status:      Open,
	}
}

func (t *Task) MarkAsDone() {
	t.Status = Done
}

func (t *Task) MarkAsInProgress() {
	t.Status = InProgress
}

func (t *Task) MarkAsOpen() {
	t.Status = Open
}
