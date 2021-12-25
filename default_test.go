package trace

import (
	"os"
)

func ExampleDefaultTracer() { //nolint
	Level(INFO)
	Writer(os.Stdout)
	Template(nil)

	Errorf("example - error")
	Infof("example - info")
	Warnf("example - warn")
	Debugf("example - debug")
	Tracef("example - trace")
	// Output:
	// example - error
	// example - info
	// example - warn
}
