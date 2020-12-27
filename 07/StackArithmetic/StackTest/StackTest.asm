
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// eq
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_1
D;JEQ
@SP
A=M-1
M=0
@IS_NOT_1
0;JMP
(IS_1)
@SP
A=M-1
M=-1
(IS_NOT_1)

// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1

// eq
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_2
D;JEQ
@SP
A=M-1
M=0
@IS_NOT_2
0;JMP
(IS_2)
@SP
A=M-1
M=-1
(IS_NOT_2)

// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// eq
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_3
D;JEQ
@SP
A=M-1
M=0
@IS_NOT_3
0;JMP
(IS_3)
@SP
A=M-1
M=-1
(IS_NOT_3)

// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_4
D;JGT
@SP
A=M-1
M=0
@IS_NOT_4
0;JMP
(IS_4)
@SP
A=M-1
M=-1
(IS_NOT_4)

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_5
D;JGT
@SP
A=M-1
M=0
@IS_NOT_5
0;JMP
(IS_5)
@SP
A=M-1
M=-1
(IS_NOT_5)

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_6
D;JGT
@SP
A=M-1
M=0
@IS_NOT_6
0;JMP
(IS_6)
@SP
A=M-1
M=-1
(IS_NOT_6)

// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// gt
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_7
D;JLT
@SP
A=M-1
M=0
@IS_NOT_7
0;JMP
(IS_7)
@SP
A=M-1
M=-1
(IS_NOT_7)

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

// gt
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_8
D;JLT
@SP
A=M-1
M=0
@IS_NOT_8
0;JMP
(IS_8)
@SP
A=M-1
M=-1
(IS_NOT_8)

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// gt
@SP
AM=M-1
D=M
@SP
A=M-1
D=D-M
@IS_9
D;JLT
@SP
A=M-1
M=0
@IS_NOT_9
0;JMP
(IS_9)
@SP
A=M-1
M=-1
(IS_NOT_9)

// push constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 53
@53
D=A
@SP
A=M
M=D
@SP
M=M+1

// add
@SP
AM=M-1
D=M
@SP
A=M-1
M=D+M

// push constant 112
@112
D=A
@SP
A=M
M=D
@SP
M=M+1

// sub
@SP
AM=M-1
D=M
@SP
A=M-1
M=M-D

// neg
@SP
A=M-1
M=-M

// and
@SP
AM=M-1
D=M
@SP
A=M-1
M=D&M

// push constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1

// or
@SP
AM=M-1
D=M
@SP
A=M-1
M=D|M

// not
@SP
A=M-1
M=!M
