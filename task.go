package miniq

import "time"

type TaskStatus uint8

const (
	TaskStatusPending   TaskStatus = iota // just created, waiting to be processed
	TaskStatusRunning                     // processing by a worker
	TaskStatusCompleted                   // process failed
	TaskStatusFailed                      // process completed
	TaskStatusDelete                      // task is being deleted
)

type Task struct {
	ID           string
	Data         []byte
	CreationDate time.Time
	Channel      string
}
