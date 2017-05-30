package loads

import "jvm_go/v1.0/instructions/base"
import "jvm_go/v1.0/rtda"

//
type ILOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.index))
}

//
type ILOAD_0 struct{ base.NoOperandsInstruction }

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(0))
}

//
type ILOAD_1 struct{ base.NoOperandsInstruction }

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, uint(1))
}

//
type ILOAD_2 struct{ base.NoOperandsInstruction }

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, uint(2))
}

//
type ILOAD_3 struct{ base.NoOperandsInstruction }

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, uint(3))
}

//
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
