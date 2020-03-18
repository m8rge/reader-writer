package closer_chan

import "fmt"

type writer struct {
	out    chan int
	closer chan struct{}
}

func (w *writer) Init() <-chan int {
	w.closer = make(chan struct{})
	w.out = make(chan int)

	return w.out
}

func (w *writer) Write() {
	defer fmt.Println("write successfully")
	for i := 1; i <= 3; i++ {
		select {
		case w.out <- i:
		case <-w.closer:
		}
	}
}

func (w *writer) Close() {
	close(w.closer)
	close(w.out)
}

type reader struct{}

func (r *reader) Read(input <-chan int) {
	defer fmt.Println("read successfully")
	for elm := range input {
		fmt.Printf("%d\n", elm)
	}
}
