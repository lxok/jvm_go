package references

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
	"container/heap"
)

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()

	if arrRef == nil {
		frame.Thread().ThrowNPE()
		return
	}

	arrLen := heap.ArrayLength(arrRef)
	stack.PushInt(arrLen)
}
