package trace

import (
	"fmt"
	"time"
)

// show the message to io.Writer without check the log level
func (tracer *Tracer) logf(msg string, args ...interface{}) (err error) {
	buff := fmt.Sprintf(msg, args...)
	switch tracer.tmpl {
	case nil:
		_, err = tracer.w.Write([]byte(buff))
	default:
		ctx := CallerContext(tracer.skip_stacks)

		ctx.Msg = buff
		ctx.Now = time.Now()
		err = tracer.tmpl.Execute(tracer.w, ctx)
	}

	if err == nil {
		// add newline
		tracer.w.Write([]byte{'\n'}) //nolint
	}
	return
}

// show the message to io.Writer without check the log level
func (tracer *Tracer) Logf(msg string, args ...interface{}) (err error) {
	err = tracer.logf(msg, args...)
	return
}

// show the error message with logLevel=ERROR
func (tracer *Tracer) Errorf(msg string, args ...interface{}) {
	if tracer.level >= ERROR {
		// show the log
		tracer.logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=WARN
func (tracer *Tracer) Warnf(msg string, args ...interface{}) {
	if tracer.level >= WARN {
		// show the log
		tracer.logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=INFO
func (tracer *Tracer) Infof(msg string, args ...interface{}) {
	if tracer.level >= INFO {
		// show the log
		tracer.logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=DEBUG
func (tracer *Tracer) Debugf(msg string, args ...interface{}) {
	if tracer.level >= DEBUG {
		// show the log
		tracer.logf(msg, args...) //nolint
	}
}

// show the error message with logLevel=TRACE
func (tracer *Tracer) Tracef(msg string, args ...interface{}) {
	if tracer.level >= TRACE {
		// show the log
		tracer.logf(msg, args...) //nolint
	}
}
