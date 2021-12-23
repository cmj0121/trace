package trace

import (
	"io"
	"text/template"
)

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

// the template of the log message
func (tracer *Tracer) Template(tmpl *template.Template) *Tracer {
	tracer.tmpl = tmpl
	return tracer
}
