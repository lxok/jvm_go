package control

import "jvm_go/v1.0/instructions/base"
import "jvm_go/v1.0/rtda"

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
