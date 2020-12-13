package glow

type ConditionallyDefinition struct {
	flow    *FlowDefinition
	handler Handler
}

func (flow *FlowDefinition) Conditionally(handler Handler) *ConditionallyDefinition {
	return &ConditionallyDefinition{flow, handler}
}

func (conditionally *ConditionallyDefinition) If(condition bool) *FlowDefinition {
	flow := conditionally.flow

	if condition {
		flow.handlers = append(flow.handlers, conditionally.handler)
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
	wrappedHandler := func(subject interface{}) {
		if f(subject) == condition {
			conditionally.handler(subject)
		}
	}

	flow := conditionally.flow
	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}
