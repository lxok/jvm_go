package main

import "fmt"
import "v1.0/classfile"
import "v1.0/instructions"
import "v1.0/instructions/base"
import "v1.0/rtda"

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.OperandStack("OperandStack:&v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []code) {
	frame := thread.PopFrame()
	reader := *base.BytecodeReader{}
	for {
		//
		pc := frame.NextPC()
		reader.Reset(bytecode, pc)

		//decode
		opcode := reader.ReadUint8()
		inst := instrutions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
