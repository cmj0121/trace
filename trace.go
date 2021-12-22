package trace

import (
	"fmt"
	"io"
	"os"
)

type Level uint

const (
	ERROR Level = iota
	WARN
	INFO
	DEBUG
	TRACE
)

// the trace instance
type Tracer struct {
	w io.Writer
}

// create tracer with default settings
func New() *Tracer {
	return &Tracer{
		// set STDERR as the defualt writer
		w: os.Stderr,
	}
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
