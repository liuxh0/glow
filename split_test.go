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

	Describe("Split()", func() {
		const handlerNumber = 2

		var (
			returnValue   *glow.FlowDefinition
			handlers      []glow.HandlerFunc
			handlerCalled [handlerNumber]bool
		)

		BeforeEach(func() {
			for i := range handlerCalled {
				handlerCalled[i] = false
			}

			handlers = make([]glow.HandlerFunc, 0, handlerNumber)
			for i := 0; i < handlerNumber; i++ {
				index := i
				handler := func(subject interface{}) interface{} {
					handlerCalled[index] = true
					return subject
				}

				handlers = append(handlers, handler)
			}
		})

		JustBeforeEach(func() {
			returnValue = flow.Split(handlers...)
		})

		It("returns the receiver pointer", func() {
			Expect(returnValue).To(Equal(flow))
		})

		It("does not call any handler", func() {
			for i, v := range handlerCalled {
				Expect(v).To(BeFalse(), "Handler %d should not be callled", i)
			}
		})

		Context("calling Run()", func() {
			var followingHandlerCalledTimes int

			BeforeEach(func() {
				followingHandlerCalledTimes = 0
			})

			JustBeforeEach(func() {
				followingHandler := func(subject interface{}) interface{} {
					followingHandlerCalledTimes++
					return subject
				}

				flow.Do(followingHandler)
				flow.Run(nil)
			})

			It("all handlers are called", func() {
				for i, v := range handlerCalled {
					Expect(v).To(BeTrue(), "Handler %d should be callled", i)
				}
			})

			It("the following handler is called multiple times", func() {
				Expect(followingHandlerCalledTimes).To(Equal(handlerNumber))
			})
		})
	})
})
