package token

const (
	ILLEGAL uint8 = iota // ~
	EOF                  // \n
	IDENT                // name

	SEMICOLON //;
	COLON     // :

	PLUS     //+
	MINUS    // -
	SLASH    // /
	ASTERISK //*
	EQL      // ==
	NQL      // !==

	keyword_beg
	FUNCTION // func
	LET      // let
	RETURN   // return
	TRUE     // true
	FALSE    // false
	IF       // if
	ELSE     // else
	keyword_end
)
