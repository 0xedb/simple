package lexer

import ( 
)

type Lexer struct {
	pos, offset int
	input       string
	ch          byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	// move pos and offset
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.offset >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.offset]
	}

	l.pos = l.offset
	l.offset++
}

func (l *Lexer) NextToken() {
	


	// switch tok.

	// switch tok.
}
