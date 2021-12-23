package trace

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// the caller context
type Context struct {
	// the caller meta
	File string
	Func string
	Line int
}

// get the caller context
func CallerContext(skip int) (ctx *Context) {
	pc := make([]uintptr, 1)

	if n := runtime.Callers(skip, pc); n > 0 {
		frames := runtime.CallersFrames(pc)
		frame, _ := frames.Next()

		ctx = &Context{
			File: filepath.Base(frame.File),
			Func: frame.Function,
			Line: frame.Line,
		}
	}

	return
}

func (ctx Context) String() (str string) {
	str = fmt.Sprintf("%v#L%03d", ctx.File, ctx.Line)
	return
}
