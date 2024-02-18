package hw06pipelineexecution

import "fmt"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	out := in

	for _, stage := range stages {
		doneSig := wrap(out, done)
		out = stage(doneSig)
	}
	return out
}

func wrap(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case data, ok := <-in:
				if !ok {
					return
				}
				fmt.Println(data)
				out <- data
			}
		}
	}()
	return out
}
