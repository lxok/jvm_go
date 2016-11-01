package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndexc      uint16
	descriptorIndex uint16
	attributes      []AttributesInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:          cp,
		accessFlags: reader.readUint16(),
		nameIndexc: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(reader, cp)
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndexc)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
