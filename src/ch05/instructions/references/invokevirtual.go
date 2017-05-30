package references

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
	"container/heap"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
	kMethodRef   *heap.ConstantMethodref
	argSlotCount uint
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	if self.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		self.kMethodRef = cp.GetConstant(self.Index).(*heap.ConstantMethodref)
		self.argSlotCount = self.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(self.argSlotCount)
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	method := self.kMethodRef.GetVirtualMethod(ref)
	frame.Thread().InvokeMethod(method)
}
