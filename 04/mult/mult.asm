// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

// Multiplication is pretty simple, really.
// All you need to do is add R0 to R0, R1 times.
// We can write a loop with an iterator to accomplish this.

// As an optimization to make sure we run the loop as few times as possible,
// we can compare R0 and R1 to see which is smaller and should
// therefore be used as the stop condition.

  @i
  M=0     // set the initial value of i to 0
  @R2
  M=0     // set the initial value of R2 to 0

  @R0     // find out if R0 is smaller than R1
  D=M
  @R1
  D=D-M
  @R0_IS_SMALLER
  D;JLE
  @R1_IS_SMALLER
  D;JGE

(R0_IS_SMALLER) // set stop condition and adder if R0 is smaller
  @R0
  D=M
  @stop
  M=D
  @R1
  D=M
  @adder
  M=D
  @LOOP
  0;JMP

(R1_IS_SMALLER) // set stop condition and adder if R1 is smaller
  @R1
  D=M
  @stop
  M=D
  @R0
  D=M
  @adder
  M=D

(LOOP)    // start the loop
  @i
  D=M
  @stop
  D=D-M
  @END
  D;JEQ   // end the program if iterator == stop

  @i
  M=M+1   // add i to the iterator

  @R2
  D=M
  @adder
  D=D+M
  @R2
  M=D     // add the adder to the running total in R2

  @LOOP
  0;JMP

(END)
  @END
  0;JMP
