package main

import (
	"errors"
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
		if ct == -1 {
			continue
		}
		// initial condition: first run sets prevType to type of first char

		// If l, r parenthesis: keep track to check mismatch
		//isParen := false
		if ct == lparen {
			paren++
		}
		if ct == rparen {
			paren--
		}

		if idx > 0 {
			// If is two-part operator
			if (prevType == 2) && (ct == 2) {
				// If is inconsistent with prior part: error
				if input[idx-1] != byte(cp) {
					return nil, errors.New("bad operator")
				}
			}
			// If type is different:
			if prevType != ct || (paren > 0 && prevType >= lparen) {
				// Append to ret
				ret = append(ret, token)
				token = ""
			}
		}

		prevType = ct

		token += string(cp)

	}
	// Check l, r paren count

	if paren > 0 {
		return nil, errors.New("mismatched paren")
	}

	// Append final result
	ret = append(ret, token)

	return ret, nil
}

func inRangeInc(low, high, test int) bool {
	// Checks if test is in range of low and high, inclusive of both
	return (low <= test) && (test <= high)
}

func charType(char rune) TokenType {
	// Checks the type of the character.
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

func charIsNumber(char rune) bool {
	// Checks if character in possible number categories
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

// Magic numbers for number type
type NumberType int

const (
	bin NumberType = iota // iota starts with 0
	dec
	hex
)

func isNumber(token string) (bool, error) {
	// Checks whether a token is a number
	// Takes only lowercased strings

	// ignores empty strings
	if len(token) <= 0 {
		return false, errors.New("No length")
	}

	// if is x: return false so that value can be subbed in
	if token == "x" {
		return false, nil
	}

	// default assumed type: decimal
	nType := dec

	// set actual number type and trim prefix
	if strings.HasPrefix(token, "0b") {
		// case: binary
		nType = bin
		token = strings.TrimPrefix(token, "0b")
	} else if strings.HasPrefix(token, "0x") {
		// case: hex
		nType = hex
		token = strings.TrimPrefix(token, "0x")
	}

	for _, cp := range token {
		// iterate over characters and expand 'ok' characters based on number type
		ok := inRangeInc('0', '1', int(cp))
		if nType > bin {
			ok = ok || inRangeInc('1', '9', int(cp))
		}
		if nType == hex {
			ok = ok || inRangeInc('a', 'f', int(cp))
		}
		if !ok {
			return false, nil
		}

	}
	return true, nil
}

func greatestPrecedence(opStack []string, token string) bool {
	/*
		Check precedence and associativity, return whether
		top of stack operator has greater precedence or
		has equal precedence to token and token is left associative
	*/
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
	// implements shunting yard algorithm
	ret := make([]int, 0)
	operatorStack := make([]string, 0)
	var operator string
	// iterate through tokens
	for i := 0; i < len(tokens); i++ {
		// read token
		token := tokens[i]
		token = strings.ToLower(token)

		// perform isNum call
		isNum, err := isNumber(token)
		if err != nil {
			return -1, err
		}
		if isNum {
			// case: number
			iConv, err := strconv.ParseInt(token, 0, 64)
			if err != nil {
				return -1, errors.New("Bad int conversion")
			}
			// add to output stack
			ret = append(ret, int(iConv))

		} else if token == "x" {
			// case: previous result
			// add to output stack
			ret = append(ret, prevRes)

		} else if token == "!" {
			// case: function
			// add to top of operator stack
			operatorStack = append([]string{token}, operatorStack...)

		} else if token != "(" && token != ")" {
			// case: operator
			for len(operatorStack) >= 1 && (greatestPrecedence(operatorStack, token) && operatorStack[0] != "(") {
				//while there's a top operator, it's greater precedence than token, and it's not left paren:
				//remove top operator and apply to end of return
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = operatorEval(ret, operator)
			}
			// add current token to top of operator stack
			operatorStack = append([]string{token}, operatorStack...)

		} else if token == "(" {
			// case: lparen
			// add to top of operator stack
			operatorStack = append([]string{token}, operatorStack...)

		} else if token == ")" {
			// case: rparen
			for operatorStack[0] != "(" {
				// while top operator isn't left paren:
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				// apply operator to end of return
				ret = operatorEval(ret, operator)
				// in case parenthesis isn't only remaining element, mismatched paren
				if (operatorStack[0] != "(") && (len(operatorStack) == 1) {
					return -1, errors.New("Mismatched parenthesis")
				}
			}
			if operatorStack[0] == "(" {
				// if top operator is paren, discard top
				operatorStack = operatorStack[1:]
			}
			if len(operatorStack) > 0 && operatorStack[0] == "!" {
				// if top operator is a function, apply to end of return
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = operatorEval(ret, operator)
			}

		}
	}
	if len(operatorStack) > 0 {
		// if still operators left in operator stack:
		if (operatorStack[0] == "(") || (operatorStack[0] == ")") {
			// if top is a parenthesis, error
			return -1, errors.New("Mismatched parenthesis")
		}
		for len(operatorStack) > 0 {
			// otherwise, apply remaining operators to end of return
			operator, operatorStack = operatorStack[0], operatorStack[1:]
			ret = operatorEval(ret, operator)
		}
	}

	// return final result
	return ret[0], nil
}

func pow(base, pow int) int {
	// power of base ^ pow
	// start ret at base
	ret := base
	if pow == 0 {
		// shortcut for power of 0
		return 1
	}
	for x := 1; x < pow; x++ {
		// multiply ret by base pow times
		ret *= base
	}
	return ret
}

func blankGen(in int) int {
	// provide a "blank" (fully 1s) int with same length as in
	final := 0
	// define max power:
	mpower := 0
	// while mpower provides a result less than input, increment mpower
	for pow(2, mpower) <= in {
		// add a new 1 onto final
		final += (1 << mpower)
		mpower++
	}
	return final
}

func operatorEval(numStack []int, operator string) []int {
	/*
		Evaluates operator for last 2 elements of numStack
		Returns a copy of the numStack with the evaluated
		elements replaced with the result.
	*/
	var arg1, arg2 int
	sLen := len(numStack)
	arg1 = numStack[sLen-1]
	if len(numStack) >= 2 {
		// Only measure 2nd arg if greater than 2 length,
		// to allow for cases where there's 1 arg with a function applied
		arg2 = numStack[sLen-2]
	}
	// arg2 will be before arg1.
	switch operator {
	// evaluate each of the operators as necessary
	case "+":
		return append(numStack[0:sLen-2], arg2+arg1)
	case "*":
		return append(numStack[0:sLen-2], arg2*arg1)
	case "-":
		return append(numStack[0:sLen-2], arg2-arg1)
	case "^":
		return append(numStack[0:sLen-2], pow(arg2, arg1))
	case "/":
		return append(numStack[0:sLen-2], arg2/arg1)
	case "&&":
		return append(numStack[0:sLen-2], arg2&arg1)
	case "||":
		return append(numStack[0:sLen-2], arg2|arg1)
	case ">>":
		return append(numStack[0:sLen-2], arg2>>arg1)
	case "<<":
		return append(numStack[0:sLen-2], arg2<<arg1)
	case "!":
		blank := blankGen(arg1)
		// taking XOR of arg1 with a "blank" int of all 1s
		// is equivalent to bitwise not
		return append(numStack[0:sLen-1], blank^arg1)
	}
	return nil
}
