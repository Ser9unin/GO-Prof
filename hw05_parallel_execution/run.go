package hw05parallelexecution

import (
	"errors"
	"fmt"
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
	doneChan := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(tasksChan <-chan Task, errChan chan<- error, doneChan <-chan struct{}) {
			for task := range tasksChan {
				err := task()
				select {
				case errChan <- err:
					fmt.Println("write ", err)
				case <-doneChan:
					return
				}
			}
			wg.Done()
		}(tasksChan, errChan, doneChan)
	}

	wg.Add(1)
	go func(tasksChan chan<- Task, doneChan <-chan struct{}) {
		defer close(tasksChan)
		for _, task := range tasks {
			select {
			case tasksChan <- task:
				fmt.Println("write to taskChan")
			case <-doneChan:
				return
			}
		}
		wg.Done()
	}(tasksChan, doneChan)

	for j := 0; j < len(tasks); j++ {
		t, ok := <-errChan
		if !ok {
			break
		}
		if t != nil {
			mCheck++
		}
		if mCheck >= m {
			err = ErrErrorsLimitExceeded
			close(doneChan)
			return err
		}
		if j == len(tasks)-1 {
			close(errChan)
		}
	}

	wg.Wait()

	return err
}
