package constants

import "v1.0/instructions/base"
import "v1.0/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {}
