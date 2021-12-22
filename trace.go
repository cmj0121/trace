package trace

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type logLevel uint

const (
	ERROR logLevel = iota
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

	// the log level of the tracer
	level logLevel

	// name of the logger
	name string
}

// create tracer with default settings
func New() *Tracer {
	return &Tracer{
		// set STDERR as the defualt writer
		w: os.Stderr,
		// set ERROR as the defualt level
		level: ERROR,
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

// change the log level
func (tracer *Tracer) Level(level logLevel) *Tracer {
	tracer.level = level
	return tracer
}

// show the message to io.Writer without check the log level
func (tracer *Tracer) Logf(msg string, args ...interface{}) (n int, err error) {
	buff := fmt.Sprintf(msg, args...) + "\n"
	n, err = tracer.w.Write([]byte(buff))
	return
}

// show the error message with logLevel=ERROR
func (tracer *Tracer) Errorf(msg string, args ...interface{}) {
	if tracer.level >= ERROR {
		// show the log
		tracer.Logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=WARN
func (tracer *Tracer) Warnf(msg string, args ...interface{}) {
	if tracer.level >= WARN {
		// show the log
		tracer.Logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=INFO
func (tracer *Tracer) Infof(msg string, args ...interface{}) {
	if tracer.level >= INFO {
		// show the log
		tracer.Logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=DEBUG
func (tracer *Tracer) Debugf(msg string, args ...interface{}) {
	if tracer.level >= DEBUG {
		// show the log
		tracer.Logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=TRACE
func (tracer *Tracer) Tracef(msg string, args ...interface{}) {
	if tracer.level >= TRACE {
		// show the log
		tracer.Logf(msg, args...) //nolint
	}
}
