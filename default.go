package trace

import (
	"io"
	"text/template"
)

var (
	DEFAULT_NAME = "default"
	// the default tracer
	default_tracer = GetTracer(DEFAULT_NAME)
)

func init() {
	// change the skip stacks to 5 for wrapper
	default_tracer.skip_stacks = 5
}

// change the default tracer's writer
func Writer(w io.Writer) {
	default_tracer.Writer(w) //nolint
}

// change the default tracer's log level
func Level(level logLevel) {
	default_tracer.Level(level) //nolint
}

// the template of the log message
func Template(tmpl *template.Template) {
	default_tracer.Template(tmpl) //nolint
}

// show the message to io.Writer without check the log level
func Logf(msg string, args ...interface{}) (err error) {
	err = default_tracer.Logf(msg, args...)
	return
}

// show the error message with logLevel=ERROR
func Errorf(msg string, args ...interface{}) {
	default_tracer.Errorf(msg, args...)
}

// show the error message with logLevel=WARN
func Warnf(msg string, args ...interface{}) {
	default_tracer.Warnf(msg, args...)
}

// show the error message with logLevel=INFO
func Infof(msg string, args ...interface{}) {
	default_tracer.Infof(msg, args...)
}

// show the error message with logLevel=DEBUG
func Debugf(msg string, args ...interface{}) {
	default_tracer.Debugf(msg, args...)
}

// show the error message with logLevel=TRACE
func Tracef(msg string, args ...interface{}) {
	default_tracer.Tracef(msg, args...)
}
