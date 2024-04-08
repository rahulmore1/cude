package worker

import (
	"cude/task"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I Will collect Stats")
}

func(w *Worker) RunTask( {
	fmt.Println("I will run a task")
}

func(w*Worker) StartTask(){
	fmt.Println("I will start a task")
}

func(w*Worker) StopTask(){
	fmt.Println("I will stop a task")
}
