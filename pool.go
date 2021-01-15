// poll
// author: baoqiang
// time: 2021/1/14 11:33 下午
package gopool

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	TASK_CHAN_SIZE = 100000
)

type pool struct {
	poolSize   int
	resultChan chan *Result
	taskChan   chan *Task
	workers    []*Worker
}

type HandleFunc func(interface{}) *Result

func NewPool(poolSize int) *pool {
	p := &pool{
		poolSize:   poolSize,
		resultChan: make(chan *Result, TASK_CHAN_SIZE),
		taskChan:   make(chan *Task, TASK_CHAN_SIZE),
	}

	for i := 0; i < poolSize; i++ {
		w := NewWorker(i+1, p.taskChan, p.resultChan)
		p.workers = append(p.workers, w)

		go func() {
			w.Run()
		}()
	}

	return p
}

func (p *pool) AddTask(t *Task) {
	p.taskChan <- t
}

func (p *pool) Wait() {
	for {
		if len(p.taskChan) == 0 {
			break
		}
		log.Debugf("pool waiting, len=%v", len(p.taskChan))
		time.Sleep(time.Millisecond * 10)
	}

	for _, w := range p.workers {
		w.Close()
	}
	close(p.taskChan)
	close(p.resultChan)
}

func (p *pool) ResultChan() <-chan *Result {
	return p.resultChan
}
