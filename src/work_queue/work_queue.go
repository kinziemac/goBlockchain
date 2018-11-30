package work_queue

import (
	// "fmt"
)

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
	WorkerNum uint
	StopRequests chan uint
	// StopRequestsNum uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	queue := new(WorkQueue)
	//sets a buffer

	queue.Jobs = make(chan Worker, maxJobs)
	queue.Results = make(chan interface{}, maxJobs)
	queue.WorkerNum = nWorkers
	queue.StopRequests = make(chan uint, queue.WorkerNum)

	// TODO: initialize struct; start nWorkers workers as goroutines
	for i := 0; i < int(nWorkers); i++ {
		go queue.worker()
	}

	return queue
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {

	//is already a go routine
	for true {
		job := <- queue.Jobs
		queue.Results <- job.Run()

		// if queue.StopRequestsNum > 0 {
		if len(queue.StopRequests) > 0 {
			//can I just remove a stop Request like this?
			<- queue.StopRequests
			// value := <- queue.StopRequests
			// fmt.Println("Removing stop request", value)
			// fmt.Println("Queue Results", len(queue.Results))
			return
		}

	}
	// TODO: Listen on the .Jobs channel for incoming tasks. For each task...
	// TODO: run tasks by calling .Run(),
	// TODO: send the return value back on Results channel.
	// TODO: Exit (return) when .Jobs is closed.
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO: close .Jobs and remove all remaining jobs from the channel.
	// closeChannel := make(chan Worker)

	//might not be able to do WorkerNum
	// for len(queue.Jobs) > 0 {
		//not sure how to empty out Jobs yet
		//hopefully emptying to new channel counts

	for i := uint(0); i < queue.WorkerNum; i++ {
			queue.StopRequests <- i
			// queue.StopRequestsNum++
	}

	//removes all remaining jobs from .Job channel
	for len(queue.Jobs) > 0 {
		<- queue.Jobs
	}

	close(queue.Jobs)
}