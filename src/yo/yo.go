package yo

import (
	"fmt"

	u "github.com/cbot918/liby/util"
)

const (
	tok_eof = -1
	tok_def = -2
	tok_extern = -3
	tok_identifier = -4
	tok_number = -5
)

type Yo struct {
	Program 				string
	LastChar 				string
	Counter 				int32
	IdentifierStr 	string
	NumVal					int32
}

func New(program string) *Yo{
	y := new(Yo)
	y.Program = program
	y.LastChar = " "  // byte_number of space
	return y
}

func (y *Yo) Run(){
	// fmt.Println("welcome to Yo Language !!")
	y.GetToken()
	// u.Logg(y.LastChar)
}


func (y *Yo) GetToken() string {
	for y.isSpace(y.LastChar) { 
		y.GetChar() 
	}

	if (y.isAlpha(y.LastChar)) {
		y.IdentifierStr = y.LastChar
		for y.isAlNum( y.GetChar() ) {
			y.IdentifierStr += y.LastChar
			// fmt.Printf("%s", y.IdentifierStr )
		}
		u.Logg(y.IdentifierStr)
		if y.IdentifierStr == "extern" { return "tok_extern" }
		if y.IdentifierStr == "def" { return "tok_def" }
		return "tok_identifier"
	}

	if (y.isDigit(y.LastChar)) {
		u.Logg("in isDigit")
		var number string
		number += y.LastChar
		for y.isDigit(y.LastChar){
			number += y.GetChar()
		}
		u.Logg(number)
		return "tok_number"
	}

	if (y.LastChar == "#") {
		for  {
			// fmt.Println(y.LastChar)
			y.GetChar()
			if y.LastChar == "\n" {
				y.GetChar()
				return y.GetToken()
			}
		}

	}

	return "token not define"
}

func (y *Yo) GetChar() string {
	y.LastChar = string(y.Program[y.Counter])
	y.Counter ++
	return y.LastChar
}

func (y *Yo) isSpace( char string ) bool{
	return char == " "
}

func (y *Yo) isAlpha( char string) bool {
	return char >= "a" && char <= "z"
}

func (y *Yo) isAlNum( char string) bool {
	return char >= "a" && char <= "z" || char >= "A" && char <= "Z" || char >= "0" && char <= "9"  
}

func (y *Yo) isDigit( char string) bool {
	return char >= "0" && char <= "9"
}

func (y *Yo) ListString() {
	for pos,char := range y.Program {
		fmt.Printf("%d: %q\n",pos, char )
	}
}
