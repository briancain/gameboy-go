package cpu

// This file serves as documentation for the organization of CPU instructions
// across multiple files. The instructions are grouped by type in the following files:
//
// - basic_instructions.go:
//   - NOP, HALT, STOP
//   - DI, EI (Disable/Enable Interrupts)
//   - 8-bit register load instructions (LD r,r', LD r,n)
//
// - load_16bit_instructions.go:
//   - 16-bit load instructions (LD rr,d16)
//   - Stack operations (PUSH, POP)
//
// - arithmetic_8bit_instructions.go:
//   - 8-bit arithmetic operations (INC r, DEC r)
//   - ADD A,r and ADC A,r instructions
//
// - arithmetic_16bit_instructions.go:
//   - 16-bit arithmetic operations (INC rr, DEC rr)
//   - ADD HL,rr and ADD SP,r8 instructions
//
// - logical_instructions.go:
//   - Logical operations (AND, OR, XOR)
//   - Immediate versions of these operations
//
// - compare_instructions.go:
//   - Comparison instructions (CP r)
//   - Subtraction instructions (SUB r, SBC A,r)
//
// - control_flow_instructions.go:
//   - Jump instructions (JP, JR)
//   - Call and return instructions (CALL, RET)
//   - Restart instructions (RST)
//
// - special_instructions.go:
//   - Special instructions (DAA, CPL, CCF, SCF)
//   - Memory access instructions
//
// - rotation_instructions.go:
//   - Standard rotation instructions (RLCA, RRCA, RLA, RRA)
//
// - cb_bit_test_instructions.go:
//   - BIT b,r instructions (Test bit b of register r)
//
// - cb_bit_reset_instructions.go:
//   - RES b,r instructions (Reset bit b of register r)
//
// - cb_bit_set_instructions.go:
//   - SET b,r instructions (Set bit b of register r)
//
// This organization makes the codebase more maintainable by grouping
// related instructions together in separate files.
