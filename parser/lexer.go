package parser

import (
	"bufio"
	"os"
)

// lexical analysis/scanning
/*
	The lexical analyzer reads the stream of characters making up the source
	program and groups the characters into meaningful sequences called lexemes.
*/

/*
For each lexeme, the lexical analyzer produces as output a token of the form:
`<token-name, attribute value>`
that it passes on to the subsequent phase, syntax analysis
*/

type token struct {
	Value string
}

func LexicalScan(file *os.File) []token {
	var bufScanner = bufio.NewScanner(file)
	var tokenList []token
	bufScanner.Split(bufio.ScanRunes)
	for bufScanner.Scan() {
		tokenList = append(tokenList, token{Value: bufScanner.Text()})
	}

	return tokenList
}
