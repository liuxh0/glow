package glow

// Split splits the flow. Each handler will be called in parallel and each branch will flow independently.
func (flow *FlowDefinition) Split(handlers ...HandlerFunc) *FlowDefinition {
	wrappedHandler := func(subject interface{}) []interface{} {
		subjects := make([]interface{}, 0, len(handlers))

		for _, handler := range handlers {
			subjects = append(subjects, handler(subject))
		}

		return subjects
	}

	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}
