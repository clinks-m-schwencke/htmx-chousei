package types

type Task struct {
	TaskId       int
	Done         bool
	Reviewed     bool
	Title        string
	Author       int
	AuthorName   string
	Assigned     int
	AssignedName string
	DueOn        string
}
