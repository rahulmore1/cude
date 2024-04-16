package main

import (
	"cude/manager"
	"cude/node"
	"cude/task"
	"cude/worker"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	var t = task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}
	var te = task.TaskEvent{
		ID:    uuid.New(),
		State: task.Pending,
		Task:  t}

	fmt.Println("task: %v \n", t)
	fmt.Println("task event: %v \n", te)

	w := worker.Worker{
		Name:  "Worker-1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	fmt.Println("Worker: %v \n", w)
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string][]*task.Task),
		EventDb: make(map[string][]*task.TaskEvent),
		Workers: []string{w.Name},
	}

	fmt.Println("Manager: %v \n", m)
	m.SelectWorker()
	m.UpdateTask()
	m.SendWork()

	n := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.1.1.",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}
	fmt.Println("Node: %v \n", n)
}
