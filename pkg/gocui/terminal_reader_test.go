package gocui

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/vt"
	"github.com/stretchr/testify/assert"
)

type delayedReadTerminal struct {
	vt.MockTerm
	holdRead     atomic.Bool
	readPaused   chan struct{}
	releaseRead  chan struct{}
	readReturned chan struct{}
	stopped      chan struct{}
	stopOnce     sync.Once
}

func (term *delayedReadTerminal) Read(data []byte) (int, error) {
	n, err := term.MockTerm.Read(data)
	if term.holdRead.CompareAndSwap(true, false) {
		close(term.readPaused)
		<-term.releaseRead
		close(term.readReturned)
	}
	return n, err
}

func (term *delayedReadTerminal) Stop() error {
	err := term.MockTerm.Stop()
	term.stopOnce.Do(func() { close(term.stopped) })
	return err
}

func TestSuspendWaitsForTerminalRead(t *testing.T) {
	term := &delayedReadTerminal{
		MockTerm:     vt.NewMockTerm(vt.MockOptSize{X: 80, Y: 24}),
		readPaused:   make(chan struct{}),
		releaseRead:  make(chan struct{}),
		readReturned: make(chan struct{}),
		stopped:      make(chan struct{}),
	}
	screen, err := tcell.NewTerminfoScreenFromTty(term)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.NoError(t, screen.Init()) {
		return
	}
	t.Cleanup(screen.Fini)
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(term.releaseRead) }) }
	t.Cleanup(release)

	term.holdRead.Store(true)
	assert.NoError(t, term.Drain())
	select {
	case <-term.readPaused:
	case <-time.After(time.Second):
		t.Fatal("terminal reader did not pause")
	}
	suspended := make(chan error, 1)
	go func() { suspended <- screen.Suspend() }()
	select {
	case <-term.stopped:
	case <-time.After(time.Second):
		t.Fatal("suspend did not stop the terminal")
	}

	returnedBeforeRead := false
	select {
	case err := <-suspended:
		assert.NoError(t, err)
		returnedBeforeRead = true
	case <-time.After(time.Second):
	}
	release()
	if !returnedBeforeRead {
		select {
		case err := <-suspended:
			assert.NoError(t, err)
		case <-time.After(time.Second):
			t.Fatal("suspend did not finish after the reader returned")
		}
	}
	<-term.readReturned
	assert.False(t, returnedBeforeRead)
	assert.NoError(t, screen.Resume())
}
