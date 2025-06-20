package task

import "time"

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
	StatusDeleted   TaskStatus = "deleted"
)

type Task struct {
	ID             string     `json:"id"`
	Status         TaskStatus `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	Duration       string     `json:"duration"`
	Result         string     `json:"result,omitempty"`
	ProcessingTime int        `json:"processing_time"`
}
