package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc uint16
	localVariable uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc: reader.readUint16(),
			localVariable: reader.readUint16(),
		}
	}
}
