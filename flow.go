package glow

// FlowDefinition describes a flow. It is used to register handlers and run the flow.
type FlowDefinition struct {
	handlers []handlerFunc
}

type handlerFunc func(subject interface{}) bool

type HandlerFunc func(subject interface{})

// FilterFunc determines whether the subject should be filtered out.
type FilterFunc func(subject interface{}) bool

// NewFlow returns a new FlowDefinition.
func NewFlow() *FlowDefinition {
	return &FlowDefinition{
		handlers: make([]handlerFunc, 0),
	}
}

// Do registers a handler in the flow that will be simply executed.
func (flow *FlowDefinition) Do(handler HandlerFunc) *FlowDefinition {
	wrappedHandler := func(subject interface{}) bool {
		handler(subject)
		return true
	}

	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}

func (flow *FlowDefinition) Filter(f FilterFunc) *FlowDefinition {
	wrappedHandler := func(subject interface{}) bool {
		return f(subject)
	}

	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}

// Run executes all registered handlers.
func (flow *FlowDefinition) Run(subject interface{}) {
	for _, handler := range flow.handlers {
		shouldContinue := handler(subject)
		if !shouldContinue {
			break
		}
	}
}
