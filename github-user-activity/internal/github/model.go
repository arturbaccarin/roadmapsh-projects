package github

type UserEvent struct {
	Type string `json:"type"`
	Repo Repo   `json:"repo"`
}

type Repo struct {
	Name string `json:"name"`
}
