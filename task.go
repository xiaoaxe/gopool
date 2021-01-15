// 任务
// author: baoqiang
// time: 2021/1/15 12:22 上午
package gopool

type Task struct {
	f   HandleFunc
	arg interface{}
}
