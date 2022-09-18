package core

type TaskStatus interface {
	GetFaild() TaskStatus
	GetFinished() TaskStatus
}

type DefaultTaskState int

const (
	FAILED DefaultTaskState = -1
	NEW    DefaultTaskState = iota
	CHECKING
	WORKING
	FINISHED
)

func (state DefaultTaskState) GetFaild() TaskStatus {
	return FAILED
}

func (state DefaultTaskState) GetFinished() TaskStatus {
	return FINISHED
}
