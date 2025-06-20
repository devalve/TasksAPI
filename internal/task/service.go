package task

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Service struct {
	mu    sync.Mutex
	tasks map[string]*Task
}

func NewService() *Service {
	return &Service{
		tasks: make(map[string]*Task),
	}
}

func (s *Service) CreateTask() *Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	taskID := strconv.Itoa(rand.Intn(1000000))
	task := &Task{
		ID:             taskID,
		Status:         StatusPending,
		CreatedAt:      time.Now(),
		ProcessingTime: rand.Intn(120) + 180, // 3-5 минут
	}

	s.tasks[taskID] = task

	go s.processTask(task)

	return task
}

func (s *Service) processTask(task *Task) {
	s.mu.Lock()
	task.Status = StatusRunning
	s.mu.Unlock()

	start := time.Now()
	time.Sleep(time.Duration(task.ProcessingTime) * time.Second)

	s.mu.Lock()
	defer s.mu.Unlock()

	if task.Status != StatusDeleted {
		task.Status = StatusCompleted
		task.Result = "Task completed successfully"
		task.Duration = time.Since(start).String()
	}
}

func (s *Service) GetTask(id string) (*Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *Service) DeleteTask(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return false
	}

	task.Status = StatusDeleted
	delete(s.tasks, id)
	return true
}
