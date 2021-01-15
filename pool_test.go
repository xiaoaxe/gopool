// pool test
// author: baoqiang
// time: 2021/1/14 11:34 下午
package gopool

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	POOL_SIZE = 3
	TASK_SIZE = 10
)

func TimeSleep(src interface{}) *Result {
	//log.Infof("start: %v", src)
	start := time.Now()
	time.Sleep(time.Duration(rand.Int31n(1000000)) * time.Microsecond * 1)
	//log.Infof("end: %v, elapsed=%s", src, time.Since(start))
	return &Result{
		Val: fmt.Sprintf("ret=%v, elapsed=%v", src, time.Since(start)),
		Err: nil,
	}
}

func TestPool(t *testing.T) {
	log.Info("run start")

	p := NewPool(POOL_SIZE)

	for idx := range [TASK_SIZE]int64{} {
		p.AddTask(&Task{
			f:   TimeSleep,
			arg: idx,
		})
	}

	p.Wait()

	for res := range p.ResultChan() {
		log.Infof("Got val=%v, err=%v", res.Val, res.Err)
	}

	log.Info("run complete")
}
