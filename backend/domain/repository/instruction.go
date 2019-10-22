package repository

func NewInstruction(persistence Persistence) Instruction {
	return &instruction{persistence: persistence}
}

type Instruction interface {
}

type instruction struct {
	persistence Persistence
}
