package constants

import "jvm_go/v1.0/instructions/base"
import "jvm_go/v1.0/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
