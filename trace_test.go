package trace

import (
	"bytes"
	"os"
	"testing"
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
	tracer.Writer(&buff)

	msg := "This is the test example message"
	nbytes, err := tracer.Logf(msg)

	if err != nil {
		// write log fail
		t.Fatalf("cannot write log: %v", err)
	}

	if nbytes != len(msg)+1 {
		// check number of bytes write
		t.Errorf("expect write %d+1 bytes: %d", len(msg), nbytes)
	}

	if buff.String() != msg+"\n" {
		// check the log message, with extra newline
		t.Errorf("expect write log %#v + '\n': %#v", msg, buff.String())
	}
}

func Example() {
	tracer := New()

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
