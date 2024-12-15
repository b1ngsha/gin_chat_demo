package utils

import "sync"

var Safety ThreadSafety

type ThreadSafety struct {
	mutex sync.Mutex
}

func (ts *ThreadSafety) ThreadSafetyDo(x func() interface{}) interface{} {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	return x()
}
