package dto

type UserEvent struct {
	Type string
	Repo Repo
}

type Repo struct {
	Name string
}
