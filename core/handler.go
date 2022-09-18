package core

import (
	"context"

	"github.com/sirupsen/logrus"
)

// TaskHandler Task状态机
type TaskHandler struct {
	TaskType         string
	StageTotal       int
	Info             TaskInfo
	CurSatus         TaskStatus
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
	for h.CurSatus != h.CurSatus.GetFaild() || h.CurSatus != h.CurSatus.GetFinished() {
		stage := h.StageTransferMap[h.CurSatus]
		h.CurSatus = stage.Do(ctx, &h.Info)
		err := h.StagePostProcess(ctx, h.CurSatus, h.Info.ErrMsg)
		if err != nil {
			logrus.Errorln(err)
		}
	}
}
