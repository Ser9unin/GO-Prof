package hw05parallelexecution

import (
	"errors"
	"fmt"
	"runtime"
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
		i := i
		wg.Add(1)
		go func(tasksChan <-chan Task, errChan chan<- error, doneChan <-chan struct{}) {
			defer wg.Done()
			fmt.Println("Worker started", i)
			for {
				task, ok := <-tasksChan
				if !ok {
					return
				}
				err := task()
				select {
				case errChan <- err:
					fmt.Println("read ", err)
				case <-doneChan:
					fmt.Println("Worker stopped", i)
					return
				}
			}
		}(tasksChan, errChan, doneChan)
	}

	wg.Add(1)
	go func(tasksChan chan<- Task, tasks []Task, doneChan <-chan struct{}) {
		defer wg.Done()
		defer close(tasksChan)
		// defer close(errChan)
		for i, task := range tasks {
			select {
			case tasksChan <- task:
				fmt.Println("write to taskChan", i)
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
				fmt.Println("Goroutines running", runtime.NumGoroutine())
				fmt.Println("Done channel closed")

				break
			}
		}
	}

	close(doneChan)
	wg.Wait()
	fmt.Println("Goroutines running", runtime.NumGoroutine())
	return err
}
