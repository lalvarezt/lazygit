package tasks

import (
	"bytes"
	"io"
	"os/exec"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jesseduffield/lazygit/pkg/gocui"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/stretchr/testify/assert"
)

type progressWriter struct{ lines atomic.Int32 }

func (w *progressWriter) Write(p []byte) (int, error) {
	w.lines.Add(int32(bytes.Count(p, []byte{'\n'})))
	return len(p), nil
}

func TestReadToEndPaintsBeforeEOF(t *testing.T) {
	writer := &progressWriter{}
	painted := make(chan int32, 16)
	firstPaint := make(chan struct{}, 1)
	noop := func() {}
	manager := NewViewBufferManager(
		utils.NewDummyLog(), writer, noop,
		func() {
			select {
			case painted <- writer.lines.Load():
			default:
			}
		},
		noop, noop, noop,
		func() { firstPaint <- struct{}{} },
		func() gocui.Task { return gocui.NewFakeTask() },
		func(f func()) error { f(); return nil },
	)
	defer manager.Close()
	reader := &BlockingLineReader{linesToYield: 3, blocked: make(chan struct{}), unblock: make(chan struct{})}
	defer close(reader.unblock)
	start := func() (Cmd, io.Reader) {
		return ExecCmd{Cmd: exec.Command("unused")}, reader
	}
	assert.NoError(t, manager.NewTask(manager.NewCmdTask(start, nil, LinesToRead{Total: 1, InitialRefreshAfter: 1}, nil), "cmd"))
	<-firstPaint
	finished := make(chan struct{})
	manager.ReadToEnd(func() { close(finished) })
	<-reader.blocked
	timeout := time.NewTimer(time.Second)
	defer timeout.Stop()
	for {
		select {
		case count := <-painted:
			if count >= 3 {
				select {
				case <-finished:
					t.Fatal("read-to-end finished before the producer reached EOF")
				default:
				}
				return
			}
		case <-timeout.C:
			t.Fatal("new lines were not painted while waiting for EOF")
		}
	}
}
