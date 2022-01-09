package trace

import (
	"fmt"
	"io"
	"os"
	"strings"
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

func LevelFromStr(str string) logLevel {
	levels := map[string]logLevel{
		"ERROR":   ERROR,
		"WARN":    WARN,
		"WARNING": WARN,
		"INFO":    INFO,
		"DEBUG":   DEBUG,
		"TRACE":   TRACE,
	}

	if level, ok := levels[strings.ToUpper(str)]; ok {
		return level
	}

	return ERROR
}

var (
	// named of the tracer
	named_tracer = map[string]*Tracer{}
	// global lock of the tracer, used when create new tracer
	tracer_lock = sync.Mutex{}
)

// any kind of pre-defined log template
var (
	// only show the raw message
	TMPL_RAW_MESSAGE = template.Must(template.New("tmpl-raw").Parse("{{ .Msg }}"))
	// show the message with timestamp and caller info
	TMPL_DEFAULT = template.Must(template.New("tmpl").Parse("[{{ .RFC3339 }}] {{ .File }}#L{{ .Line }} - {{ .Msg }}"))
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
	tracer := &Tracer{
		// set STDERR as the defualt writer
		w: os.Stderr,
		// set ERROR as the defualt level
		level: ERROR,
		// skip stacks of caller
		skip_stacks: 4,
		// set the defualt template
		tmpl: TMPL_DEFAULT,
	}

	tracer.prologue()
	return tracer
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

	tracer.prologue()
	return tracer
}

// the necessary setup when call Tracer
func (tracer *Tracer) prologue() {
	level := os.Getenv(strings.ToUpper(fmt.Sprintf("%v_LOG_LEVEL", tracer.name)))
	tracer.level = LevelFromStr(level)
}
