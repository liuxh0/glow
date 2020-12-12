package glow

type FlowDefinition struct {
	handlers []Handler
}

func NewFlow() *FlowDefinition {
	return &FlowDefinition{
		handlers: make([]Handler, 0),
	}
}

func (flow *FlowDefinition) Run(subject interface{}) {
	for _, handler := range flow.handlers {
		handler(subject)
	}
}

func (flow *FlowDefinition) Do(handler Handler) *FlowDefinition {
	flow.handlers = append(flow.handlers, handler)
	return flow
}
