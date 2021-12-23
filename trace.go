package trace

import (
	"io"
	"os"
	"sync"
	"text/template"
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

	// the template of the logger
	tmpl *template.Template

	// skip stacks of caller
	skip_stacks int
}

// create tracer with default settings
func New() *Tracer {
	return &Tracer{
		// set STDERR as the defualt writer
		w: os.Stderr,
		// set ERROR as the defualt level
		level: ERROR,
		// skip stacks of caller
		skip_stacks: 4,
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
