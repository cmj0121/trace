package trace

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Level uint

const (
	ERROR Level = iota
	WARN
	INFO
	DEBUG
	TRACE
)

var (
	// named of the tracer
	named_tracer = map[string]*Tracer{}
	// global lock of the tracer, used when create new tracer
	tracer_lock = sync.Mutex{}
)

// the trace instance
type Tracer struct {
	// the destination writer of the logger
	w io.Writer
	// name of the logger
	name string
}

// create tracer with default settings
func New() *Tracer {
	return &Tracer{
		// set STDERR as the defualt writer
		w: os.Stderr,
	}
}

// create or get the named tracer
func GetTracer(name string) *Tracer {
	tracer_lock.Lock()
	defer tracer_lock.Unlock()

	if tracer, ok := named_tracer[name]; ok {
		// found the named tracer
		return tracer
	}

	tracer := New()
	tracer.name = name
	named_tracer[name] = tracer
	return tracer
}

// override the writer
func (tracer *Tracer) Writer(w io.Writer) *Tracer {
	tracer.w = w
	return tracer
}

// show the message to io.Writer
func (tracer *Tracer) Logf(msg string, args ...interface{}) (n int, err error) {
	buff := fmt.Sprintf(msg, args...) + "\n"
	n, err = tracer.w.Write([]byte(buff))
	return
}
