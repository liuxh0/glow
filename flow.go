package glow

// FlowDefinition describes a flow. It is used to register handlers and run the flow.
type FlowDefinition struct {
	handlers []Handler
}

type Handler func(subject interface{})

// NewFlow returns a new FlowDefinition.
func NewFlow() *FlowDefinition {
	return &FlowDefinition{
		handlers: make([]Handler, 0),
	}
}

// Do registers a handler in the flow that will be simply executed.
func (flow *FlowDefinition) Do(handler Handler) *FlowDefinition {
	flow.handlers = append(flow.handlers, handler)
	return flow
}

// Run executes all registered handlers.
func (flow *FlowDefinition) Run(subject interface{}) {
	for _, handler := range flow.handlers {
		handler(subject)
	}
}
