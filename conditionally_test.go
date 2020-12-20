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

	Describe("Conditionally()", func() {
		var conditionally *glow.ConditionallyDefinition

		BeforeEach(func() {
			conditionally = flow.Conditionally(nil)
		})

		It("returns a non-nil pointer", func() {
			Expect(conditionally).NotTo(BeNil())
		})
	})
})

var _ = Describe("ConditionallyDefinition", func() {
	var (
		flow          *glow.FlowDefinition
		conditionally *glow.ConditionallyDefinition
		handlerCalled bool
	)

	BeforeEach(func() {
		handlerCalled = false
		handler := func(subject interface{}) interface{} {
			handlerCalled = true
			return subject
		}

		flow = glow.NewFlow()
		conditionally = flow.Conditionally(handler)
	})

	Describe("If()", func() {
		var (
			condition    bool
			returnedFlow *glow.FlowDefinition
		)

		BeforeEach(func() {
			condition = false
		})

		JustBeforeEach(func() {
			returnedFlow = conditionally.If(condition)
			flow.Run(nil)
		})

		When("condition is true", func() {
			BeforeEach(func() {
				condition = true
			})

			It("returns the same FlowDefinition pointer", func() {
				Expect(returnedFlow).To(Equal(flow))
			})

			Specify("the handler is called when running the flow", func() {
				Expect(handlerCalled).To(BeTrue())
			})
		})

		When("condition is false", func() {
			BeforeEach(func() {
				condition = false
			})

			It("returns the same FlowDefinition pointer", func() {
				Expect(returnedFlow).To(Equal(flow))
			})

			Specify("the handler is not called when running the flow", func() {
				Expect(handlerCalled).To(BeFalse())
			})
		})
	})

	Describe("IfNot()", func() {
		var (
			condition    bool
			returnedFlow *glow.FlowDefinition
		)

		BeforeEach(func() {
			condition = false
		})

		JustBeforeEach(func() {
			returnedFlow = conditionally.IfNot(condition)
			flow.Run(nil)
		})

		When("condition is true", func() {
			BeforeEach(func() {
				condition = true
			})

			It("returns the same FlowDefinition pointer", func() {
				Expect(returnedFlow).To(Equal(flow))
			})

			Specify("the handler is not called when running the flow", func() {
				Expect(handlerCalled).To(BeFalse())
			})
		})

		When("condition is false", func() {
			BeforeEach(func() {
				condition = false
			})

			It("returns the same FlowDefinition pointer", func() {
				Expect(returnedFlow).To(Equal(flow))
			})

			Specify("the handler is called when running the flow", func() {
				Expect(handlerCalled).To(BeTrue())
			})
		})
	})

	Describe("IfFuncReturnsTrue()", func() {
		var (
			conditionFuncReturn bool
			conditionFuncCalled bool
			returnedFlow        *glow.FlowDefinition
		)

		BeforeEach(func() {
			conditionFuncReturn = false
			conditionFuncCalled = false
		})

		JustBeforeEach(func() {
			conditionFunc := func(_ interface{}) bool {
				conditionFuncCalled = true
				return conditionFuncReturn
			}

			returnedFlow = conditionally.IfFuncReturnsTrue(conditionFunc)
		})

		It("does not call condition function", func() {
			Expect(conditionFuncCalled).To(BeFalse())
		})

		It("returns the same FlowDefinition pointer", func() {
			Expect(returnedFlow).To(Equal(flow))
		})

		Context("running the flow", func() {
			JustBeforeEach(func() {
				flow.Run(nil)
			})

			Specify("the condition function is called", func() {
				Expect(conditionFuncCalled).To(BeTrue())
			})

			When("condition function returns true", func() {
				BeforeEach(func() {
					conditionFuncReturn = true
				})

				Specify("the handler is called", func() {
					Expect(handlerCalled).To(BeTrue())
				})
			})

			When("condition function returns false", func() {
				BeforeEach(func() {
					conditionFuncReturn = false
				})

				Specify("the handler is not called", func() {
					Expect(handlerCalled).To(BeFalse())
				})
			})
		})
	})

	Describe("IfFuncReturnsFalse()", func() {
		var (
			conditionFuncReturn bool
			conditionFuncCalled bool
			returnedFlow        *glow.FlowDefinition
		)

		BeforeEach(func() {
			conditionFuncReturn = false
			conditionFuncCalled = false
		})

		JustBeforeEach(func() {
			conditionFunc := func(_ interface{}) bool {
				conditionFuncCalled = true
				return conditionFuncReturn
			}

			returnedFlow = conditionally.IfFuncReturnsFalse(conditionFunc)
		})

		It("does not call condition function", func() {
			Expect(conditionFuncCalled).To(BeFalse())
		})

		It("returns the same FlowDefinition pointer", func() {
			Expect(returnedFlow).To(Equal(flow))
		})

		Context("running the flow", func() {
			JustBeforeEach(func() {
				flow.Run(nil)
			})

			Specify("the condition function is called", func() {
				Expect(conditionFuncCalled).To(BeTrue())
			})

			When("condition function returns true", func() {
				BeforeEach(func() {
					conditionFuncReturn = true
				})

				Specify("the handler is not called", func() {
					Expect(handlerCalled).To(BeFalse())
				})
			})

			When("condition function returns false", func() {
				BeforeEach(func() {
					conditionFuncReturn = false
				})

				Specify("the handler is called", func() {
					Expect(handlerCalled).To(BeTrue())
				})
			})
		})
	})
})
