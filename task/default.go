package task

import (
	"context"
	"fmt"
	"state-machine/core"
)

type DefaultTaskNew struct{}
type DefaultTaskChecking struct{}
type DefaultTaskWorking struct{}
type DefaultTaskFinished struct{}

var DefaultTaskType = "default task"
var (
	defaultTaskNew         DefaultTaskNew
	defaultTaskChecking    DefaultTaskChecking
	defaultTaskWorking     DefaultTaskWorking
	DefaultTaskTransferMap = core.TaskTransferMap
)

func init() {
	transfermap := map[core.TaskStatus]core.TaskStage{
		core.NEW:      &defaultTaskNew,
		core.CHECKING: &defaultTaskChecking,
		core.WORKING:  &defaultTaskWorking,
	}
	DefaultTaskTransferMap[DefaultTaskType] = transfermap
}

func (s *DefaultTaskNew) Do(ctx context.Context, task *core.TaskInfo) core.TaskStatus {
	fmt.Printf("task start")
	return core.CHECKING
}

func (s *DefaultTaskChecking) Do(ctx context.Context, task *core.TaskInfo) core.TaskStatus {
	// TODO do some checking
	
	return core.WORKING
}

func (s *DefaultTaskWorking) Do(ctx context.Context, task *core.TaskInfo) core.TaskStatus {
	// TODO do some working

	return core.FINISHED
}

func DefaultTaskPostProcess(ctx context.Context, status core.TaskStatus, errMsg string) {
	// TODO do some post process
}
