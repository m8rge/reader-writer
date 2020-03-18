package closer_chan

import (
	"fmt"
	"sync"
)

type writer struct {
	out      chan int
	closer   chan struct{}
	outMutex sync.Mutex
}

func (w *writer) Init() <-chan int {
	w.closer = make(chan struct{})
	w.out = make(chan int)

	return w.out
}

func (w *writer) Write() {
	for i := 1; i <= 3; i++ {
		select {
		case <-w.closer:
			return
		default:
		}

		w.outMutex.Lock()
		select {
		case w.out <- i:
			w.outMutex.Unlock()
		case <-w.closer:
			w.outMutex.Unlock()
			return
		}
	}

	return
}

func (w *writer) Close() {
	close(w.closer)
	w.outMutex.Lock()
	close(w.out)
	w.outMutex.Unlock()
}

type reader struct{}

func (r *reader) Read(input <-chan int) {
	for elm := range input {
		fmt.Printf("%d\n", elm)
	}
}
