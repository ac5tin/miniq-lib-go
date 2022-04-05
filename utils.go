package miniq

import (
	"github.com/ac5tin/miniq-lib-go/proto"
)

func ProtoTask2Task(t *proto.Task) *Task {
	return &Task{
		ID:           t.Id,
		Data:         t.Data,
		CreationDate: t.CreationDate.AsTime(),
		Channel:      t.Channel,
	}
}
