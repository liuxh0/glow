package glow

// FlowDefinition describes a flow. It is used to register handlers and run the flow.
type FlowDefinition struct {
	handlers []handlerFunc
	subjects []interface{}
}

type handlerFunc func(subject interface{}) []interface{}

type HandlerFunc func(subject interface{}) interface{}

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
	wrappedHandler := func(subject interface{}) []interface{} {
		return []interface{}{handler(subject)}
	}

	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}

func (flow *FlowDefinition) Filter(f FilterFunc) *FlowDefinition {
	wrappedHandler := func(subject interface{}) []interface{} {
		if f(subject) == true {
			return []interface{}{subject}
		}

		return nil
	}

	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}

// Run executes all registered handlers.
func (flow *FlowDefinition) Run(subject interface{}) {
	flow.subjects = []interface{}{subject}

	for _, handler := range flow.handlers {
		nextSubjects := make([]interface{}, 0, len(flow.subjects))

		for _, subject := range flow.subjects {
			subjects := handler(subject)
			if subjects != nil {
				nextSubjects = append(nextSubjects, subjects...)
			}
		}

		flow.subjects = nextSubjects
		if len(flow.subjects) == 0 {
			return
		}
	}
}
