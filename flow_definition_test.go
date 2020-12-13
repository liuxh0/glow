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

		When("calling Run() after it", func() {
			var handlerCalled = false

			BeforeEach(func() {
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
