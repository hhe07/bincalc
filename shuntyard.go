package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Magic numbers for Tokenizer
type TokenType int

const (
	num TokenType = iota
	funct 
	operator 
	lparen 
	rparen 
)

// Tokenizer: separates input string into its component parts
func Tokenizer(input string) ([]string, error) {
	ret := make([]string, 0)
	token := ""
	var prevType TokenType // 0: num, 1: func, 2: operator, 3/4: l/r paren
	paren := 0
	/* 
		Iterates over every codepoint,
		determines where to separate tokens
		by comparing to previous character
		type.
	*/
	for idx, cp := range input {
		ct := charType(cp)
		// initial condition: first run sets prevType to type of first char
		if idx == 0{
			prevType = ct
		}
		// If is two-part operator
		if (prevType == 2) && (ct == 2) {
			// If is inconsistent with prior part: error
			if input[idx - 1] != byte(cp) {
				return nil, errors.New("bad operator")
			}
		}
		
		// If l, r parenthesis: keep track to check mismatch
		//isParen := false
		if ct == lparen{
			paren++
		}
		if ct == rparen{
			paren--
		}

		// If type is different:
		if (prevType != ct || paren > 0) && !(idx == 0){
			if prevType != -1{
				// If previous wasn't just a space:
				// Append to ret
				ret = append(ret, token)
			}
			token = ""
		}
		prevType = ct

		token += string(cp)

	}
	// Check l, r paren count
	
	if paren > 0{
		return nil, errors.New("mismatched paren")
	}
	
	// Append final result
	ret = append(ret, token)
	
	return ret, nil
}

func inRangeInc(low, high, test int) bool {
	// Inclusive inRange
	return (low <= test) && (test <= high)
}

func charType(char rune) TokenType {
	if charIsNumber(char) {
		return num
	} else if charIsFunc(char) {
		return funct
	} else if charIsOperator(char) {
		return operator
	} else if charIsLParen(char) {
		return lparen
	} else if charIsRParen(char) {
		return rparen
	}
	return -1
}

// Magic numbers for number type
type NumberType int

const (
	bin NumberType = iota
	dec
	hex
)

func charIsNumber(char rune) bool {
	ic := int(char)
	// return: is number               is hex part                   is x
	return (inRangeInc('0', '9', ic) || (inRangeInc('a', 'f', ic) || char == 'x'))
}

func charIsFunc(char rune) bool {
	// Checks if character in functions list
	functions := []rune{'!'}
	// Iterate over all possible functions
	for _, elem := range functions {
		if char == elem {
			return true
		}
	}
	return false
}

func isNumber(token string) (bool, error) {
	// Takes only lowercased strings
	if len(token) <= 0 {
		return false, errors.New("No length")
	}

	if token == "x"{
		return false, nil
	}

	nType := dec
	
	if strings.HasPrefix(token, "0b"){
		nType = bin
		token = strings.TrimPrefix(token, "0b")
	} else if strings.HasPrefix(token, "0x"){
		nType = hex
		token = strings.TrimPrefix(token, "0x")
	}
	
	for _, cp := range token {
		// CP is the unicode codepoint
		ok := inRangeInc('0', '1', int(cp))
		if nType > bin{
			ok = ok || inRangeInc('1', '9', int(cp))
		}
		if nType == hex{
			ok = ok || inRangeInc('a', 'f', int(cp))
		}
		if !ok{
			return false, nil
		}

	}
	return true, nil

}


func charIsOperator(char rune) bool {
	// Checks if character in operators list
	operators := []rune{'^', '*', '/', '+', '-', '&', '|', '<', '>'}
	// Iterate over all possible operators
	for _, elem := range operators {
		if char == elem {
			return true
		}
	}
	return false
}

func charIsLParen(char rune) bool {
	// Checks if char is left parenthesis
	return char == '('
}

func charIsRParen(char rune) bool {
	// Checks if char is right parenthesis
	return char == ')'
}

func greatestPrecedence(opStack []string, token string) bool {
	// Check precedence and associativity, return whether
	// top of stack has greatest precedence or
	// has equal precedence to rest and is left associative
	opPrec := map[string]int{
		"^":  4,
		"*":  3,
		"/":  3,
		"+":  2,
		"-":  2,
		"&&": 0,
		"||": 0,
		"<<": 0,
		">>": 0,
	}
	rightAs := []string{"^", "<<", ">>"}
	topPrec := opPrec[token]
	currPrec := opPrec[opStack[0]]

	isRightAs := false

	// Checks if operator is right associative
	for _, r := range rightAs {
		if r == token {
			isRightAs = true
		}
	}

	return (currPrec > topPrec) || ((currPrec == topPrec) && !isRightAs)
}

func ShuntYard(tokens []string, prevRes int) (int, error) {
	// todo: soft failure?
	ret := make([]int, 0)
	operatorStack := make([]string, 0)
	var operator string
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		token = strings.ToLower(token)

		// Handle number
		isNum, err := isNumber(token)
		if err != nil {
			fmt.Println(err)
		}

		if isNum {
			// case: number
			iConv, err := strconv.ParseInt(token, 0, 64)
			if err != nil {
				return -1, errors.New("Bad int conversion")
			}
			ret = append(ret, int(iConv))
			
		} else if token == "x"{
			// case: previous result
			ret = append(ret, prevRes)
		} else if token == "!" {
			// case: function
			operatorStack = append([]string{token}, operatorStack...)
			
		} else if token != "(" && token != ")" {
			// case: operator
			for len(operatorStack) > 1 && (greatestPrecedence(operatorStack, token) && operatorStack[0] != "(") {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = operatorEval(ret, operator)
			}
			operatorStack = append([]string{token}, operatorStack...)
			
		} else if token == "(" {
			// case: lparen
			operatorStack = append([]string{token}, operatorStack...)
			
		} else if token == ")" {
			// case: rparen
			for operatorStack[0] != "(" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = operatorEval(ret, operator)
				if (operatorStack[0] != "(") && (len(operatorStack) == 1) {
					return -1, errors.New("Mismatched parenthesis")
				}
			}
			if operatorStack[0] == "(" {
				operatorStack = operatorStack[1:]
			}
			if operatorStack[0] == "!" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = operatorEval(ret, operator)
			}

		}
	}
	if len(operatorStack) > 0 {
		if (operatorStack[0] == "(") || (operatorStack[0] == ")") {
			return -1, errors.New("Mismatched parenthesis")
		}
		for len(operatorStack) > 0 {
			operator, operatorStack = operatorStack[0], operatorStack[1:] 
			ret = operatorEval(ret, operator)
		}
	}

	return ret[0], nil
}

func pow(base, pow int) int {
	ret := base
	if pow == 0{
		return 1
	}
	for x := 1; x < pow; x++ {
		ret *= base
	}
	return ret
}

func blankGen(in int) int{
	final := 0
	// define max power:
	mpower := 0
	for pow(2, mpower) <= in{
		final += (1<<mpower)
		mpower++
	}
	return final
}

func operatorEval(numStack []int, operator string) []int {
	var arg1, arg2 int
	sLen := len(numStack)
	arg1 = numStack[sLen-1]
	if len(numStack) >=2{
		arg2 = numStack[sLen-2]	
	}
	switch operator {
	case "+":
		return append(numStack[0:sLen-2], arg2 + arg1) 
	case "*":
		return append(numStack[0:sLen-2], arg2 * arg1) 
	case "-":
		return append(numStack[0:sLen-2], arg2 - arg1) 
	case "^":
		return append(numStack[0:sLen-2], pow(arg2, arg1)) 
	case "/":
		return append(numStack[0:sLen-2], arg2 / arg1) 
	case "&&":
		return append(numStack[0:sLen-2], arg2 & arg1)
	case "||":
		return append(numStack[0:sLen-2], arg2 | arg1)
	case ">>":
		return append(numStack[0:sLen-2], arg2 >> arg1)
	case "<<":
		return append(numStack[0:sLen-2], arg2 << arg1)
	case "!":
		blank := blankGen(arg1)
		return append(numStack[0:sLen-1], blank ^ arg1)
	}
	return nil
}
