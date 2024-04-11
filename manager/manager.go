package manager

import (
	"cude/task"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate Worker")
}

func (m *Manager) UpdateTask() {
	fmt.Println("I will Update the tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to the workers")
}
