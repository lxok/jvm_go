package references

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
	"container/heap"
)

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()

	cp := frame.ConstantPool()
	kClass := cp.GetConstant(self.Index).(*heap.ConstantClass)
	class := kClass.Class()

	if ref == nil {
		stack.PushInt(0)
	} else if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
