package loads

import "jvm_go/v1.0/instructions/base"
import "jvm_go/v1.0/rtda"

//
type LLOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(self.index))
}

//
type LLOAD_0 struct{ base.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, uint(0))
}

//
type LLOAD_1 struct{ base.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, uint(1))
}

//
type LLOAD_2 struct{ base.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, uint(2))
}

//
type LLOAD_3 struct{ base.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, uint(3))
}

//
func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
