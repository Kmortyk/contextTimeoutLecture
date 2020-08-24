package main

import (
	"context"
	"fmt"
	"time"
)

type Worker struct {
	idx int

	programCtx context.Context
	workerCtx  context.Context

	cancel    context.CancelFunc
	isStopped bool
}

func NewWorker(idx int, ctx context.Context) *Worker {
	workerCtx, cancel := context.WithCancel(context.Background())

	return &Worker{
		idx:        idx,
		programCtx: ctx,
		workerCtx:  workerCtx,
		cancel:     cancel,
		isStopped:  false,
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			// программа завершилась
			case <-w.programCtx.Done():
				return
			// завершили одного воркера
			case <-w.workerCtx.Done():
				return
			default:
			}

			time.Sleep(time.Second)
			fmt.Println("worker", w.idx, "processing job...")
		}
	}()
}

func (w *Worker) Stop() {
	w.cancel()
	w.isStopped = true
}

func (w *Worker) IsStopped() bool {
	return w.isStopped
}

func main() {
	const numWorkers = 5
	ctx, cancel := context.WithCancel(context.Background())
	workers := make([]*Worker, numWorkers)

	for workerIdx := 0; workerIdx < numWorkers; workerIdx++ {
		workers[workerIdx] = NewWorker(workerIdx, ctx)
		workers[workerIdx].Start()
	}

	// останавливаем второго
	workers[2].Stop()
	time.Sleep(3 * time.Second)

	// завершаем остальных
	cancel()
}
