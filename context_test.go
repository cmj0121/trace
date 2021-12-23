package trace

import (
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := CallerContext(1)
	line := 28

	if ctx == nil {
		t.Fatalf("cannot get caller context")
	}

	if ctx.File != "context.go" {
		// should be the CallerContext
		t.Errorf("expect get context.File = context.go: %v", ctx.File)
	}

	if ctx.Func != "github.com/cmj0121/trace.CallerContext" {
		// should be the github.com/cmj0121/trace.CallerContext
		t.Errorf("expect get context.Func = github.com/cmj0121/trace.CallerContext: %v", ctx.Func)
	}

	if ctx.Line != 28 {
		// should be the 20
		t.Errorf("expect get context.Line = %d: %v", line, ctx.Line)
	}

	if ctx.String() != fmt.Sprintf("context.go#L%03d", line) {
		// should be context.go#L021
		t.Errorf("expect get context = context.go#L%03d: %v", line, ctx.String())
	}
}
