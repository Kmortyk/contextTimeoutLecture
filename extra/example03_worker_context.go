package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	Пример пул воркеров, каждый хранит два контекста - один общий,
	другой - для остановки конкретного воркера.
*/

type Worker struct {
	idx int

	programCtx context.Context
	workerCtx  context.Context

	cancel context.CancelFunc

	mu        sync.RWMutex
	isStopped bool
}

func NewWorker(ctx context.Context, idx int) *Worker {
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

	w.mu.Lock()
	defer w.mu.Unlock()
	w.isStopped = true
}

func (w *Worker) IsStopped() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.isStopped
}

func main() {
	const numWorkers = 5
	ctx, cancel := context.WithCancel(context.Background())
	workers := make([]*Worker, numWorkers) // slice из nil'ов

	for workerIdx := 0; workerIdx < numWorkers; workerIdx++ {
		workers[workerIdx] = NewWorker(ctx, workerIdx)
		workers[workerIdx].Start()
	}

	// останавливаем второго
	workers[2].Stop()
	time.Sleep(10 * time.Second)

	// завершаем остальных
	cancel()
}
