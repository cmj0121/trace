package trace

import (
	"bytes"
	"os"
	"testing"
	"text/template"
)

type MockWriter struct {
}

func (mock MockWriter) Write(p []byte) (n int, err error) {
	// mock with nop
	return
}

func TestLogf(t *testing.T) {
	var buff bytes.Buffer

	tracer := New()
	tracer.Writer(&buff).Template(nil)

	msg := "This is the test example message"
	err := tracer.Logf(msg)

	if err != nil {
		// write log fail
		t.Fatalf("cannot write log: %v", err)
	}

	if buff.String() != msg+"\n" {
		// check the log message, with extra newline
		t.Errorf("expect write log %#v + '\n': %#v", msg, buff.String())
	}
}

func TestNamedTracer(t *testing.T) {
	tracer_foo := GetTracer("foo")
	tracer_dup := GetTracer("foo")
	tracer_bob := GetTracer("bob")

	if tracer_foo != tracer_dup {
		// get the diff tracer
		t.Errorf("expect get the same named tracer: %v %v", tracer_foo, tracer_dup)
	}

	if tracer_foo == tracer_bob {
		// get the same tracer
		t.Errorf("expect get the diff named tracer: %v %v", tracer_foo, tracer_bob)
	}
}

func Example() {
	tracer := New().Template(nil)

	name := "tracer"

	tracer.Writer(os.Stdout)
	tracer.Logf("example - %d", 1)     //nolint
	tracer.Logf("example - %v", name)  //nolint
	tracer.Logf("example - %#v", name) //nolint
	// Output:
	// example - 1
	// example - tracer
	// example - "tracer"
}

func ExampleLevel() {
	tracer := New().Level(INFO).Writer(os.Stdout).Template(nil)

	tracer.Errorf("example - error")
	tracer.Warnf("example - warn")
	tracer.Infof("example - info")
	tracer.Debugf("example - debug")
	tracer.Tracef("example - trace")
	// Output:
	// example - error
	// example - warn
	// example - info
}

func ExampleTemplate() {
	tmpl := template.Must(template.New("tmpl").Parse("{{ .File }}#L{{ .Line }} - {{ .Msg }}"))
	tracer := New().Writer(os.Stdout).Template(tmpl).Level(INFO)

	tmpl_2 := template.Must(template.New("tmpl").Parse("[{{ .File }}#L{{ .Line }}] - {{ .Msg }}"))
	Writer(os.Stdout)
	Template(tmpl_2)
	Level(INFO)

	tracer.Errorf("example - error")
	tracer.Warnf("example - warn")
	tracer.Infof("example - info")
	tracer.Debugf("example - debug")
	tracer.Tracef("example - trace")
	Errorf("default - error")
	Warnf("default - warn")
	Infof("default - info")
	Debugf("default - debug")
	Tracef("default - trace")
	// Output:
	// trace_test.go#L92 - example - error
	// trace_test.go#L93 - example - warn
	// trace_test.go#L94 - example - info
	// [trace_test.go#L97] - default - error
	// [trace_test.go#L98] - default - warn
	// [trace_test.go#L99] - default - info
}

func BenchmarkLogf(b *testing.B) {
	mock_writer := MockWriter{}

	tracer := New()
	tracer.Writer(mock_writer)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// fixed string with one and only one parameter
			tracer.Logf("benchmark:  %v", pb) //nolint
		}
	})
}
