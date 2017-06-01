package constants

import "v1.0/instructions/base"
import "v1.0/rtda"

//
type BIPUSH struct{ val int8 }

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}

//
type SIPUSH struct{ val int16 }

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}
