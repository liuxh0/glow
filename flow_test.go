package glow_test

import (
	"fmt"

	"github.com/liuxh0/glow"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewFlow()", func() {
	var flowDefinition *glow.FlowDefinition

	BeforeEach(func() {
		flowDefinition = glow.NewFlow()
	})

	It("returns a non-nil pointer", func() {
		Expect(flowDefinition).NotTo(BeNil())
	})
})

var _ = Describe("FlowDefinition", func() {
	var flow *glow.FlowDefinition

	BeforeEach(func() {
		flow = glow.NewFlow()
	})

	Describe("Do()", func() {
		It("returns the receiver pointer", func() {
			Expect(flow.Do(nil)).To(Equal(flow))
		})

		Context("calling Run() after it", func() {
			var handlerCalled bool

			BeforeEach(func() {
				handlerCalled = false

				handler := func(_ interface{}) {
					handlerCalled = true
				}

				flow.Do(handler).Run(nil)
			})

			Specify("the handler is called", func() {
				Expect(handlerCalled).To(BeTrue())
			})
		})
	})

	Describe("Filter()", func() {
		var (
			filterFuncReturn bool
			filterFuncCalled bool
		)

		BeforeEach(func() {
			filterFuncReturn = false
			filterFuncCalled = false
		})

		JustBeforeEach(func() {
			filterFunc := func(_ interface{}) bool {
				filterFuncCalled = true
				return filterFuncReturn
			}

			flow.Filter(filterFunc)
		})

		It("returns the receiver pointer", func() {
			Expect(flow.Filter(nil)).To(Equal(flow))
		})

		It("does not call filter function", func() {
			Expect(filterFuncCalled).To(BeFalse())
		})

		Context("registering a handler after it and then calling Run()", func() {
			var (
				handlerFuncCalled bool
			)

			BeforeEach(func() {
				handlerFuncCalled = false
			})

			JustBeforeEach(func() {
				handlerFunc := func(_ interface{}) {
					handlerFuncCalled = true
				}

				flow.Do(handlerFunc)
				flow.Run(nil)
			})

			Specify("the filter function is called", func() {
				Expect(filterFuncCalled).To(BeTrue())
			})

			When("the filter function returns true", func() {
				BeforeEach(func() {
					filterFuncReturn = true
				})

				Specify("the following handler function is called", func() {
					Expect(handlerFuncCalled).To(BeTrue())
				})
			})

			When("the filter function returns false", func() {
				BeforeEach(func() {
					filterFuncReturn = false
				})

				Specify("the following handler function is not called", func() {
					Expect(handlerFuncCalled).To(BeFalse())
				})
			})
		})
	})

	Describe("Run()", func() {
		When("no hanlder is registered", func() {
			It("doesn't panic", func() {
				Expect(func() { flow.Run(nil) }).NotTo(Panic())
			})
		})

		When("multiple handlers are registered", func() {
			const (
				handlerNum   = 3
				outputFormat = "Handler %d"
			)

			var (
				handlerOutput []string
			)

			BeforeEach(func() {
				handlerOutput = make([]string, 0, handlerNum)

				for i := 0; i < handlerNum; i++ {
					index := i
					handler := func(_ interface{}) {
						handlerOutput = append(handlerOutput, fmt.Sprintf(outputFormat, index))
					}

					flow.Do(handler)
				}

				flow.Run(nil)
			})

			It("calls each handler", func() {
				Expect(handlerOutput).To(HaveLen(handlerNum))
				for i, output := range handlerOutput {
					Expect(output).To(Equal(fmt.Sprintf(outputFormat, i)))
				}
			})
		})
	})
})
