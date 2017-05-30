package references

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
	"container/heap"
)

// Invoke a class (static) method
type INVOKE_STATIC struct {
	base.Index16Instruction
	method *heap.Method
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	if self.method == nil {
		cp := frame.Method().Class().ConstantPool()
		k := cp.GetConstant(self.Index)
		if kMethodRef, ok := k.(*heap.ConstantMethodref); ok {
			self.method = kMethodRef.StaticMethod()
		} else {
			self.method = k.(*heap.ConstantInterfaceMethodref).StaticMethod()
		}
	}

	// init class
	class := self.method.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	frame.Thread().InvokeMethod(self.method)
}
