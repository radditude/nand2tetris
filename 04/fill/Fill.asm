// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed.
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

(CHECK_KEYS)
  @SCREEN
  D=A
  @current
  M=D               // save a pointer for the start of the screen
                    // we'll use this in the loops later

  @KBD
  D=M
  @ENDARKINATOR
  D;JNE             // jump if @KEYBOARD is not zero (e.g., a key is pressed)
  @ENLIGHTENATOR
  D;JEQ             // jump if @KEYBOARD is zero


(ENDARKINATOR)
  @current
  A=M
  M=-1

  @current          // check if we're all the way through the screen
  M=M+1
  D=M
  @KBD
  D=A-D
  @CHECK_KEYS
  D;JLE

  @ENDARKINATOR
  0;JMP             // if not, move to the next address and loop again


(ENLIGHTENATOR)
  @current
  A=M
  M=0

  @current          // check if we're all the way through the screen
  M=M+1
  D=M
  @KBD
  D=A-D
  @CHECK_KEYS
  D;JEQ

  @ENLIGHTENATOR
  0;JMP             // if not, go to the next screen address
