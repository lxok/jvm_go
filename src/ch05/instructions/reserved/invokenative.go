package reserved

import (
	"ch05/instructions/base"
	"ch05/jvm/rtda"
)

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	nativeMethod := frame.Method().NativeMethod().(func(*rtda.Frame))
	nativeMethod(frame)
}
