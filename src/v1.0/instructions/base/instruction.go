package base

import "jvm_go/v1.0/rtda"

//interface
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

//abstract class
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

//abstract class
type BranchInstruction struct {
	Offset int
}

func (seft *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	seft.Offset = int(reader.ReadInt16())
}

//abstract class
type Index8Instrction struct {
	Index uint
}

func (seft *Index8Instrction) FetchOperands(reader *BytecodeReader) {
	seft.Index = uint(reader.readUint8())
}

//abstract class
type Index16Instrction struct {
	Index uint
}

func (seft *Index16Instrction) FetchOperands(reader *BytecodeReader) {
	seft.Index = uint(reader.readUint16())
}
