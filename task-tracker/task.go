package main

type Tasks map[string]Task

type status string

const (
	Open       status = "open"
	InProgress status = "in_progress"
	Done       status = "done"
)

type Task struct {
	Title  string
	Body   string
	Status status
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
