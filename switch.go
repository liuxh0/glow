package glow

// SwitchFunc returns a string to indicate which handler to run.
type SwitchFunc func(subject interface{}) string

// Switch switches the flow to one of the defined handlers, according to the string that the switch function returns.
// "_default" is reserved for specifying the default handler.
func (flow *FlowDefinition) Switch(switchFunc SwitchFunc, handlerMap map[string]HandlerFunc) *FlowDefinition {
	wrappedHandler := func(subject interface{}) []interface{} {
		key := switchFunc(subject)
		if handler, ok := handlerMap[key]; ok {
			return []interface{}{handler(subject)}
		} else if defaultHandler, ok := handlerMap["_default"]; ok {
			return []interface{}{defaultHandler(subject)}
		}

		return nil
	}

	flow.handlers = append(flow.handlers, wrappedHandler)
	return flow
}
