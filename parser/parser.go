package parser

import "fmt"

func isObject(tokenList []token) bool {
	last := len(tokenList) - 1
	if tokenList[0].Value == "{" && tokenList[last].Value == "}" {
		return true
	}

	return false
}

func Parse(tokenList []token) (string, error) {

	if len(tokenList) == 0 {
		return "", fmt.Errorf("empty file, invalid JSON")
	}

	if isObject(tokenList) {
		return "valid JSON", nil
	}
	return "", fmt.Errorf("not a valid object")
}
