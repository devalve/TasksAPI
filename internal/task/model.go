package task

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusDeleted   Status = "deleted"
)

type Task struct {
	ID             string    `json:"id"`
	Status         Status    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	Duration       string    `json:"duration"`
	Result         string    `json:"result,omitempty"`
	ProcessingTime int       `json:"processing_time"`
}
