package constants

import "jvm_go/v1.0/instructions/base"
import "jvm_go/v1.0/rtda"

//
type BIPUSH struct{ val int8 }

func (seft *BIPUSH) FetchOperands(reader *BytecodeReader) {
	seft.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(seft.val))
}

//
type SIPUSH struct{ val int16 }

func (seft *SIPUSH) FetchOperands(reader *BytecodeReader) {
	seft.val = int(reader.ReadInt16())
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(seft.val))
}
