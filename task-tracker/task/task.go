package main

import "fmt"

type status string

const (
	Open       status = "open"
	InProgress status = "in_progress"
	Done       status = "done"
)

type Tasks map[string]Task

func (t Tasks) Add(task *Task) {
	t[task.Title] = *task
}

func (t Tasks) Remove(title string) {
	delete(t, title)
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

type Task struct {
	ID          int
	Description string
	Status      status
	createdAt   string
	updatedAt   string
}

func (t Task) String() string {
	return fmt.Sprintf("Title: %s\nBody: %s\nStatus: %s\n", t.Title, t.Body, t.Status)
}

func NewTask(title, body string) Task {
	return Task{
		Title:  title,
		Body:   body,
		Status: Open,
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
