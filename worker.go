// worker
// author: baoqiang
// time: 2021/1/14 11:33 下午
package gopool

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Worker struct {
	id         int
	resultChan chan *Result
	taskChan   <-chan *Task
	exitChan   chan struct{}
}

func NewWorker(id int, taskChan chan *Task, resultChan chan *Result) *Worker {
	return &Worker{
		id:         id,
		resultChan: resultChan,
		taskChan:   taskChan,
		exitChan:   make(chan struct{}),
	}
}

func (w *Worker) Run() {
	for {
		select {
		case t := <-w.taskChan:
			log.Debugf("got arg: %v", t.arg)
			res := t.f(t.arg)
			w.resultChan <- res
		case <-w.exitChan:
			log.Infof("Task-%d exiting...", w.id)
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (w *Worker) Close() {
	w.exitChan <- struct{}{}
}
