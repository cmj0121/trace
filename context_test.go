package trace

import (
	"testing"
)

func TestContext(t *testing.T) {
	ctx := CallerContext(1)

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

	if ctx.Line != 21 {
		// should be the 20
		t.Errorf("expect get context.Line = 21: %v", ctx.Line)
	}

	if ctx.String() != "context.go#L021" {
		// should be context.go#L021
		t.Errorf("expect get context = context.go#L20: %v", ctx.String())
	}
}
