package references

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
	"container/heap"
)

// Create new object
type NEW struct {
	base.Index16Instruction
	class *heap.Class
}

func (self *NEW) Execute(frame *rtda.Frame) {
	if self.class == nil {
		cp := frame.ConstantPool()
		kClass := cp.GetConstant(self.Index).(*heap.ConstantClass)
		self.class = kClass.Class()
	}

	// init class
	if self.class.InitializationNotStarted() {
		frame.RevertNextPC() // undo new
		frame.Thread().InitClass(self.class)
		return
	}

	ref := self.class.NewObj()
	frame.OperandStack().PushRef(ref)
}
