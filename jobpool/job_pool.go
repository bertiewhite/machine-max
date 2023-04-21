package jobpool

import (
	"sync"
)

type JobPool struct {
	Done       chan interface{} // I couldn't decide between a done channel
	buffer     chan bool
	action     func() error
	errHandler func(error)
	numWorkers int
	numJobs    int
}

func NewJobPool(action func() error, errHandler func(error), numWorkers, numJobs int) *JobPool {
	return &JobPool{
		action:     action,
		errHandler: errHandler,
		numWorkers: numWorkers,
		numJobs:    numJobs,
	}
}

// Start blocks until all messages are processed. It can, and in certain scenarios should, be used asynchronously. However,
// blocking should be the default behaviour with asynchronous-ness being dealt with the caller.
func (jp *JobPool) Start() {
	jp.buffer = make(chan bool, jp.numWorkers)

	go func() {
		for i := 0; i < jp.numJobs; i++ {
			jp.buffer <- true
		}
		close(jp.buffer)
	}()

	var wg sync.WaitGroup
	for j := 0; j < jp.numWorkers; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// Have to separate these in order to enforce the closing of the channel is dealt with preferentially.
				// We could maybe start looking at contexts to handle these more gracefully. My understanding of contexts
				// suggest they are harder and may be less suited to these.
				select {
				case _, ok := <-jp.Done:
					if !ok {
						return
					}
				default:
				}

				select {
				case _, ok := <-jp.buffer:
					if !ok {
						return
					}
					err := jp.action()
					if err != nil {
						jp.errHandler(err)
					}
				default:
				}
			}
		}()
	}
	wg.Wait()
}

// I'd be tempted to include this and initialise the done channel in the start function. For now, I prefer the management
// of the done channel being outside this object.

//func (jp *JobPool) Stop() {
//	if jp.Done != nil {
//		close(jp.Done)
//	}
//}
