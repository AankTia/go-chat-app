package trace
// Tracer is the interface that describes an object caable of
// tracing event througout code
type Tracer interface {
	Trace(...interface{})
}