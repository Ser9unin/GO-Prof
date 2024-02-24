package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	wrapCh := in

	for _, stage := range stages {
		doneSig := wrap(wrapCh, done)
		wrapCh = stage(doneSig)
	}

	return wrapCh
}

func wrap(wrapCh In, done In) Out {
	dataCh := make(Bi)

	go func() {
		defer close(dataCh)
		for {
			select {
			case <-done:
				for range wrapCh {
					// do nothing, read from "in" to prevent goroutine leak
				}
				return
			case data, ok := <-wrapCh:
				if !ok {
					return
				}
				dataCh <- data
			}
		}
	}()

	return dataCh
}
