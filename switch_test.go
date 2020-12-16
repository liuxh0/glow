package glow_test

import (
	"github.com/liuxh0/glow"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FlowDefinition", func() {
	var flow *glow.FlowDefinition

	BeforeEach(func() {
		flow = glow.NewFlow()
	})

	Describe("Switch()", func() {
		var (
			returnValue          *glow.FlowDefinition
			switchFuncReturn     string
			switchFuncCalled     bool
			handlers             []string
			handlerMap           map[string]glow.HandlerFunc
			handlerFuncCalled    []string
			defaultHandlerCalled bool
		)

		BeforeEach(func() {
			switchFuncReturn = ""
			switchFuncCalled = false
			handlers = []string{"handler1", "handler2"}
			handlerMap = make(map[string]glow.HandlerFunc)
			handlerFuncCalled = make([]string, 0)
			defaultHandlerCalled = false

			for _, v := range handlers {
				handlerName := v
				handlerFunc := func(_ interface{}) {
					handlerFuncCalled = append(handlerFuncCalled, handlerName)
				}

				handlerMap[v] = handlerFunc
			}

			defaultHandler := func(_ interface{}) {
				defaultHandlerCalled = true
			}
			handlerMap["_default"] = defaultHandler
		})

		JustBeforeEach(func() {
			switchFunc := func(_ interface{}) string {
				switchFuncCalled = true
				return switchFuncReturn
			}

			returnValue = flow.Switch(switchFunc, handlerMap)
		})

		It("returns the receiver pointer", func() {
			Expect(returnValue).To(Equal(flow))
		})

		It("does not call the switch function", func() {
			Expect(switchFuncCalled).To(BeFalse())
		})

		It("does not call any handler function", func() {
			Expect(handlerFuncCalled).To(HaveLen(0))
			Expect(defaultHandlerCalled).To(BeFalse())
		})

		Context("calling Run()", func() {
			JustBeforeEach(func() {
				flow.Run(nil)
			})

			Specify("the switch function is called", func() {
				Expect(switchFuncCalled).To(BeTrue())
			})

			When("the switch function returns a string that exists in the map", func() {
				BeforeEach(func() {
					switchFuncReturn = "handler1"
					Expect(handlers).To(ContainElement(switchFuncReturn))
				})

				Specify("the corresponding handler is called", func() {
					Expect(handlerFuncCalled).To(ContainElement("handler1"))
				})

				Specify("other handlers are not called", func() {
					for _, v := range handlers {
						if v != switchFuncReturn {
							Expect(handlerFuncCalled).NotTo(ContainElement(v))
						}
					}
				})

				Specify("the default handler is not called", func() {
					Expect(defaultHandlerCalled).To(BeFalse())
				})
			})

			When("the switch function returns a string that does not exist in the map", func() {
				BeforeEach(func() {
					switchFuncReturn = "non-exist-handler"
					Expect(handlers).NotTo(ContainElement(switchFuncReturn))
				})

				When("default handler is registered", func() {
					BeforeEach(func() {
						Expect(handlerMap).To(HaveKey("_default"))
					})

					Specify("the default handler is called", func() {
						Expect(defaultHandlerCalled).To(BeTrue())
					})

					Specify("other handlers are not called", func() {
						Expect(handlerFuncCalled).To(HaveLen(0))
					})
				})

				When("default handler is not registered", func() {
					BeforeEach(func() {
						delete(handlerMap, "_default")
						Expect(handlerMap).NotTo(HaveKey("_default"))
					})

					Specify("other handlers are not called", func() {
						Expect(handlerFuncCalled).To(HaveLen(0))
					})
				})
			})
		})
	})
})
