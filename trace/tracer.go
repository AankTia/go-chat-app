package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object caable of
// tracing event througout code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *nilTracer) Trace(a ...interface{}) {}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &nilTracer{}
}
