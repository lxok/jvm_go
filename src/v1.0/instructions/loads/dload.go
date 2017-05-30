package loads

import "jvm_go/v1.0/instructions/base"
import "jvm_go/v1.0/rtda"

//
type DLOAD struct{ base.Index8Instruction }

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(self.index))
}

//
type DLOAD_0 struct{ base.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, uint(0))
}

//
type DLOAD_1 struct{ base.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, uint(1))
}

//
type DLOAD_2 struct{ base.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, uint(2))
}

//
type DLOAD_3 struct{ base.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, uint(3))
}

//
func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
