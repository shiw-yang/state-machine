package core

type TaskStatus int

const (
	FAILED TaskStatus = -1
	NEW    TaskStatus = iota
	CHECKING
	WORKING
	FINISHED
)
