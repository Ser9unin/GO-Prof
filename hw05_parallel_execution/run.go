package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var (
		wg     sync.WaitGroup
		mCheck int
		err    error
	)

	tasksChan := make(chan Task)
	errChan := make(chan error)
	defer close(errChan)

	doneChan := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(tasksChan <-chan Task, errChan chan<- error, doneChan <-chan struct{}) {
			defer wg.Done()
			for task := range tasksChan {
				err := task()
				select {
				case errChan <- err:
				case <-doneChan:
					return
				}
			}
		}(tasksChan, errChan, doneChan)
	}

	wg.Add(1)
	go func(tasksChan chan<- Task, tasks []Task, doneChan <-chan struct{}) {
		defer wg.Done()
		defer close(tasksChan)
		for _, task := range tasks {
			select {
			case tasksChan <- task:
			case <-doneChan:
				return
			}
		}
	}(tasksChan, tasks, doneChan)

	for j := 0; j < len(tasks); j++ {
		t, ok := <-errChan
		if !ok {
			break
		}
		if t != nil {
			mCheck++
			if mCheck >= m {
				err = ErrErrorsLimitExceeded
				break
			}
		}
	}

	close(doneChan)

	wg.Wait()

	return err
}
