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
		defer func() {
			close(dataCh)
			for range wrapCh {
				// read from "wrapCh" to avoid goroutine leak
			}
		}()
		for {
			select {
			case <-done:
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
