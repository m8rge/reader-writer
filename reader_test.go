package closer_chan

import (
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestWriteRead(t *testing.T) {
	defer goleak.VerifyNone(t)

	w := &writer{}
	output := w.Init()
	go w.Write()
	r := &reader{}
	go r.Read(output)

	time.Sleep(100 * time.Millisecond)

	w.Close()

	time.Sleep(100 * time.Millisecond)
}

func TestWriteNoRead(t *testing.T) {
	defer goleak.VerifyNone(t)

	w := &writer{}
	w.Init()
	go w.Write()

	time.Sleep(100 * time.Millisecond)

	w.Close()

	time.Sleep(100 * time.Millisecond)
}

func TestNoWriteRead(t *testing.T) {
	defer goleak.VerifyNone(t)

	w := &writer{}
	output := w.Init()
	r := &reader{}
	go r.Read(output)

	time.Sleep(100 * time.Millisecond)

	w.Close()

	time.Sleep(100 * time.Millisecond)
}

func TestNoWriteNoRead(t *testing.T) {
	defer goleak.VerifyNone(t)

	w := &writer{}
	w.Init()

	w.Close()

	time.Sleep(100 * time.Millisecond)
}

func TestLateWriteRead(t *testing.T) {
	defer goleak.VerifyNone(t)

	w := &writer{}
	output := w.Init()
	r := &reader{}
	go r.Read(output)

	time.Sleep(100 * time.Millisecond)

	w.Close()
	go w.Write()

	time.Sleep(100 * time.Millisecond)
}

func TestLateWriteNoRead(t *testing.T) {
	defer goleak.VerifyNone(t)

	w := &writer{}
	w.Init()

	w.Close()
	go w.Write()

	time.Sleep(100 * time.Millisecond)
}
