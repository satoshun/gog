package main

const workerPoolSize = 20

var workerPool chan chan string

type workerTask func(string)

type worker struct {
	channel chan string
	task    workerTask
}

func (w *worker) start() {
	go func() {
		for {
			workerPool <- w.channel
			select {
			case url := <-w.channel:
				w.task(url)
			}
		}
	}()
}

func makeWorkerPool(task workerTask) {
	workerPool = make(chan chan string, workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		w := worker{channel: make(chan string), task: task}
		w.start()
	}
}
