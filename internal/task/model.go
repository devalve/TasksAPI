package task

import "time"

type Status string

const (
	Pending   Status = "pending"
	Running   Status = "running"
	Completed Status = "completed"
	Deleted   Status = "deleted"
)

type Task struct {
	ID             string    `json:"id"`
	Status         Status    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	Duration       string    `json:"duration"`
	Result         string    `json:"result,omitempty"`
	ProcessingTime int       `json:"processing_time"`
}
