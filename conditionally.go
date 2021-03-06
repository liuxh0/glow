package glow

type ConditionallyDefinition struct {
	flow    *FlowDefinition
	handler HandlerFunc
}

func (flow *FlowDefinition) Conditionally(handler HandlerFunc) *ConditionallyDefinition {
	return &ConditionallyDefinition{flow, handler}
}

func (conditionally *ConditionallyDefinition) If(condition bool) *FlowDefinition {
	flow := conditionally.flow

	if condition {
		wrappedHandler := func(subject interface{}) []interface{} {
			return []interface{}{conditionally.handler(subject)}
		}

		flow.handlers = append(flow.handlers, wrappedHandler)
	}

	return flow
}

func (conditionally *ConditionallyDefinition) IfNot(condition bool) *FlowDefinition {
	return conditionally.If(!condition)
}

func (conditionally *ConditionallyDefinition) IfFuncReturnsTrue(condition func(subject interface{}) bool) *FlowDefinition {
	return conditionally.ifFunc(condition, true)
}

func (conditionally *ConditionallyDefinition) IfFuncReturnsFalse(condition func(subject interface{}) bool) *FlowDefinition {
	return conditionally.ifFunc(condition, false)
}

func (conditionally *ConditionallyDefinition) ifFunc(f func(subject interface{}) bool, condition bool) *FlowDefinition {
	wrappedHandler := func(subject interface{}) []interface{} {
		if f(subject) == condition {
			return []interface{}{conditionally.handler(subject)}
		}

		return nil
	}

	flow := conditionally.flow
	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}
