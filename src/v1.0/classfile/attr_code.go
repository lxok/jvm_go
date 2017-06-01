package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocal       uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocal = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func (self *CodeAttribute) MaxLocals() uint16 {
	return self.maxLocal
}

func (self *CodeAttribute) MaxStack() uint16 {
	return self.maxStack
}

func (self *CodeAttribute) Code() []byte {
	return self.code
}
