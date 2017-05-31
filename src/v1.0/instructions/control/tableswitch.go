package control

import "jvm_go/v1.0/instructions/base"
import "jvm_go/v1.0/rtda"

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *BytecodeReader) {
	reader.SkipPadding()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
