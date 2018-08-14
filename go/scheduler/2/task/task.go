package task

import (
	"reflect"
	"time"
)

type Task struct {
	TaskID int64
	Func   FunctionMeta
	RunAt  time.Time
	Timer  *time.Timer
	Stop   bool
}

func NewTask(function FunctionMeta, runAt time.Time) *Task {
	return &Task{
		TaskID: time.Now().UnixNano(),
		Func:   function,
		RunAt:  runAt,
		Timer:  time.NewTimer(runAt.Sub(time.Now())),
	}
}

func (task *Task) Run() {
	function := reflect.ValueOf(task.Func.function)
	// params := make([]reflect.Value, len(task.Params))
	// for i, param := range task.Params {
	// 	params[i] = reflect.ValueOf(param)
	// }
	function.Call([]reflect.Value{})
}
