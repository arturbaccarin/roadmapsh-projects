package github

type Requester interface {
	GetListEventsUser(username string) ([]UserEvent, error)
}
