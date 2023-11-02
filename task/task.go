package task

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     string
	Status      string // pending, completed, etc.
}
