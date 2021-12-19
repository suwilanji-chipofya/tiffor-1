package virtual_machine

const PC_START=0xa

type VirtualMachine struct {
    Memory [255]uint16
    Pc uint8
    State bool
}

var Regs [4]uint16

var Opcodes=map[int]string {
    0:"HALT",
    1:"ADD",
    2:"SUB",
    3:"MUL",
    4:"DIV"
    5:"MODULO"
}

func (vm *VirtualMachine) halt() {
    vm.State=false
}

func (vm *VirtualMachine) start(instructions []uint16) {
    vm.Memory=instructions
}