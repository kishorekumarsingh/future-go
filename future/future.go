package future

import (
	"sync"
	"time"
)

// Task represents an interruptable goroutine.
type Task struct {
	running    bool
	err        error
	stopChan   chan struct{}
	shouldStop bool
	lock       sync.Mutex
	result     interface{}
	wg         sync.WaitGroup
}

// S will return true if task needs to be stopped.
type S func() bool

// Running returns true if the task is running.
func (task *Task) Running() bool {
	task.lock.Lock()
	running := task.running
	task.lock.Unlock()
	return running
}

// Cancel the task if possible.
func (task *Task) Cancel() {
	task.lock.Lock()
	task.shouldStop = true
	task.lock.Unlock()
	select {
	case <-task.StopChan():
	case <-time.After(5 * time.Second):
	}
}

// Done returns true if the task was cancelled or finished executing.
func (task *Task) Done() bool {
	task.lock.Lock()
	running := task.running
	task.lock.Unlock()
	return !running
}

// Result returns the result when execution of the task is finished.
func (task *Task) Result() interface{} {
	task.lock.Lock()
	result := task.result
	task.lock.Unlock()
	return result
}

// Exception returns any exception while executing the task.
func (task *Task) Exception() error {
	task.lock.Lock()
	err := task.err
	task.lock.Unlock()
	return err
}

// StopChan gets the stop channel for this task.
func (task *Task) StopChan() <-chan struct{} {
	return task.stopChan
}

// GetSetMega executes the function in a goroutine and returns an interruptable task.
func GetSetMega(fn func(S) interface{}) *Task {
	task := &Task{
		stopChan: make(chan struct{}),
		running:  true,
	}
	task.wg.Add(1)
	go func() {
		defer task.wg.Done()
		v := fn(func() bool {
			task.lock.Lock()
			shouldStop := task.shouldStop
			task.lock.Unlock()
			return shouldStop
		})

		// When the task is stopped.
		task.lock.Lock()
		switch v := v.(type) {
		case error:
			task.err = v
		case string:
			task.result = v
		}
		task.running = false
		close(task.stopChan)
		task.lock.Unlock()
	}()
	return task
}

// Submit runs the function as a goroutine and returns an interruptable task.
func Submit(f func(url string, timeout time.Duration) interface{}, url string, timeout time.Duration) *Task {
	task := GetSetMega(func(shouldStop S) interface{} {
	out:
		for {
			var result interface{} = f(url, timeout)
			if shouldStop() {
				break out
			}
			return result
		}
		return nil
	})
	task.wg.Wait()
	return task
}
