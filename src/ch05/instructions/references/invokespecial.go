package references

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
	"container/heap"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	k := cp.GetConstant(self.Index)
	if kMethodRef, ok := k.(*heap.ConstantMethodref); ok {
		method := kMethodRef.SpecialMethod()
		frame.Thread().InvokeMethod(method)
	} else {
		method := k.(*heap.ConstantInterfaceMethodref).SpecialMethod()
		frame.Thread().InvokeMethod(method)
	}
}
