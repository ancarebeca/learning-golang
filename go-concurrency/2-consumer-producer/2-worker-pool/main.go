package main

import (
	"log"
	"runtime"
	"time"
)

type WorkerPool interface {
	Run()
	AddTask(task func()) //function to be processed to the Worker Pool.
}

type workerPool struct {
	maxWorker  int
	queuedTask chan func()
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) { // We spawn the goroutine based on maxWorker
			for task := range wp.queuedTask { // we pull the task from queuedTask channel,
				task() // execute the task
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTask <- task
}

func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker:  maxWorker,
		queuedTask: make(chan func()),
	}

	return wp
}

func main() {

	// For monitoring purpose.
	waitC := make(chan bool)
	go func() {
		for {
			log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	// Start Worker Pool.
	totalWorker := 3
	wp := NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTask := 5
	resultC := make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i + 1
		wp.AddTask(func() {
			log.Printf("[main] Starting task %d", id)
			time.Sleep(5 * time.Second)
			resultC <- result{id, id * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task %d has been finished with result %d:", res.id, res.value)
	}

	<-waitC
}

//https://medium.com/code-chasm/go-concurrency-pattern-worker-pool-a437117025b1
