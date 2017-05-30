package references

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
	"container/heap"
)

// Set field in object
type PUT_FIELD struct {
	base.Index16Instruction
	field *heap.Field
}

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.Index).(*heap.ConstantFieldref)
		self.field = kFieldRef.InstanceField()
	}

	stack := frame.OperandStack()
	val := stack.PopField(self.field.IsLongOrDouble)
	ref := stack.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	self.field.PutValue(ref, val)
}
