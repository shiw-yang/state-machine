package core

import (
	"context"

	"github.com/sirupsen/logrus"
)

var (
	TaskTransferMap = map[string]map[TaskStatus]TaskStage{}
)

// TaskHandler Task状态机
type TaskHandler struct {
	TaskType         string
	StageTotal       int
	Info             TaskInfo
	CurStatus        TaskStatus
	StageTransferMap map[TaskStatus]TaskStage
	StagePostProcess func(ctx context.Context, status TaskStatus, errMsg string) error
}

// TaskInfo 任务信息
type TaskInfo struct {
	ErrMsg string
}

type TaskStage interface {
	Do(context.Context, *TaskInfo) TaskStatus
}

// Start start transfer
func (h *TaskHandler) Start(ctx context.Context) {
	for h.IsEnd() {
		stage := h.StageTransferMap[h.CurStatus]
		h.CurStatus = stage.Do(ctx, &h.Info)
		err := h.StagePostProcess(ctx, h.CurStatus, h.Info.ErrMsg)
		if err != nil {
			logrus.Errorln(err)
		}
	}
}

func (h *TaskHandler) IsEnd() bool {
	if h.CurStatus == FAILED || h.CurStatus == FINISHED {
		return true
	}
	return false
}
