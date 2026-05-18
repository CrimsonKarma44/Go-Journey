package main

import (
	"fmt"
	"sync"
)

type Pool struct {
	maxWorker int
	workers   []*worker
	sender    chan func()

	wg    sync.WaitGroup
	mutex sync.Mutex
}

func NewPool(maxWorker int) *Pool {
	channel := make(chan func(), 100)
	workers := make([]*worker, 0, maxWorker)

	p := Pool{
		maxWorker: maxWorker,
		workers:   workers,
		sender:    channel,
	}
	for i := range maxWorker {
		Worker := &worker{
			id:       i + 1,
			receiver: channel,
			wg:       &p.wg,
		}
		p.wg.Add(1)
		go Worker.runTasks()
		workers = append(workers, Worker)
	}

	return &p
}

func (p *Pool) AddTask(task func()) error {
	select {
	case p.sender <- task:
	default:
		return fmt.Errorf("pool full")
	}
	return nil
}

func (p *Pool) Close() {
	close(p.sender)
	p.wg.Wait()
	println("All worker closed")
}

type worker struct {
	id       int
	receiver chan func()
	wg       *sync.WaitGroup
}

func (w *worker) runTasks() {
	fmt.Printf("initializing worker {%d}\n", w.id)
	defer w.wg.Done()
	for task := range w.receiver {
		fmt.Printf("Starting task from worker {%d}\n", w.id)
		task()
	}
	fmt.Printf("worker {%d}: receiver closed; exiting\n", w.id)
}
