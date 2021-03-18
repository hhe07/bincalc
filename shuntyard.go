package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Tokenizer(input string) ([]string, error) {
	ret := make([]string, 0)
	token := ""
	prevType := 0 // 0: num, 1: func, 2: operator, 3/4: l/r paren
	lParen, rParen := 0, 0
	// todo: check parenthesis count
	// todo: check for mistyped operations, esp. boolean ones
	for idx, cp := range input {
		if idx == 0 {
			prevType = charType(cp)
		}
		if (prevType == 2) && (charType(cp) == 2) {
			if input[idx - 1] != byte(cp) {
				return nil, errors.New("bad operator")
			}
		}

		if charType(cp) == 3{
			lParen++
		}
		if charType(cp) == 4{
			rParen++
		}

		if prevType != charType(cp) {
			if prevType != -1{
				ret = append(ret, token)
			}
			prevType = charType(cp)
			token = ""
		}

		token += string(cp)

	}
	if lParen != rParen{
		return nil, errors.New("mismatched paren")
	}
	ret = append(ret, token)
	return ret, nil
}

// or have tokens separated by spaces when entering args?
func inRangeNInc(low, high, test int) bool {
	// Non-inclusive inRange
	return (low < test) && (test < high)
}

func inRangeInc(low, high, test int) bool {
	// Inclusive inRange
	return (low <= test) && (test <= high)
}

func charType(char rune) int {
	if charIsNumber(char) {
		return 0
	} else if charIsFunc(char) {
		return 1
	} else if charIsOperator(char) {
		return 2
	} else if charIsLParen(char) {
		return 3
	} else if charIsRParen(char) {
		return 4
	}
	return -1
}

func charIsNumber(char rune) bool {
	// If   not a number                        not lowercase                      not x (possible for 0x)
	return (inRangeInc('0', '9', int(char)) || (inRangeInc('a', 'z', int(char)) || char == 'x'))
}

func charIsFunc(char rune) bool {
	functions := []rune{'!'}
	for _, elem := range functions {
		if char == elem {
			return true
		}
	}
	return false
}

func charIsOperator(char rune) bool {
	operators := []rune{'^', '*', '/', '+', '-', '&', '|', '<', '>'}
	for _, elem := range operators {
		if char == elem {
			return true
		}
	}
	return false
}

func charIsLParen(char rune) bool {
	return char == '('
}

func charIsRParen(char rune) bool {
	return char == ')'
}

func isNumber(token string) (bool, error) {
	// Takes only lowercased strings
	// TODO: Make the checking stronger: x only allowable on 0x cases or just as "x".
	if len(token) <= 0 {
		return false, errors.New("No length")
	}

	if token == "x"{
		return false, nil
	}

	for _, cp := range token {
		// CP is the unicode codepoint
		if !(charIsNumber(cp)) {
			return false, nil
		}
	}
	return true, nil

}

func greatestPrecedence(opStack []string, token string) bool {
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
	for _, r := range rightAs {
		if r == token {
			isRightAs = true
		}
	}
	return (currPrec > topPrec) || ((currPrec == topPrec) && !isRightAs)
}

func ShuntYard(tokens []string, prevRes int) (int, error) {
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
		// case: number
		if isNum {
			iConv, err := strconv.Atoi(token)
			if err != nil {
				return -1, errors.New("Bad int conversion")
			}
			ret = append(ret, iConv)
			// case: function
		} else if token == "x"{
			ret = append(ret, prevRes)
		} else if token == "!" {
			operatorStack = append([]string{token}, operatorStack...)
			// case: operator
		} else if token != "(" && token != ")" {
			for len(operatorStack) > 1 && (greatestPrecedence(operatorStack, token) && operatorStack[0] != "(") {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret[0:len(ret)-2], operatorEval(ret, operator)) 
			}
			operatorStack = append([]string{token}, operatorStack...)
			// case: lparen
		} else if token == "(" {
			operatorStack = append([]string{token}, operatorStack...)
			// case: rparen
		} else if token == ")" {
			for operatorStack[0] != "(" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret[0:len(ret)-2], operatorEval(ret, operator)) 
				if (operatorStack[0] != "(") && (len(operatorStack) == 1) {
					return -1, errors.New("Mismatched parenthesis")
				}
			}
			if operatorStack[0] == "(" {
				operatorStack = operatorStack[1:]
			}
			if operatorStack[0] == "!" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret[0:len(ret)-2], operatorEval(ret, operator)) 
			}

		}
	}
	if len(operatorStack) > 0 {
		if (operatorStack[0] == "(") || (operatorStack[0] == ")") {
			return -1, errors.New("Mismatched parenthesis")
		}
		for len(operatorStack) > 0 {
			operator, operatorStack = operatorStack[0], operatorStack[1:] 
			ret = append(ret[0:len(ret)-2], operatorEval(ret, operator)) 
		}
	}

	return ret[0], nil
}

func pow(base, pow int) int {
	ret := base
	for x := 1; x < pow; x++ {
		ret *= base
	}
	return ret
}

func operatorEval(numStack []int, operator string) int {
	if len(numStack) < 2 {
		errors.New("bad numstack")
	}
	var arg1, arg2 int
	arg1 = numStack[len(numStack)-1]
	arg2 = numStack[len(numStack)-2]
	switch operator {
	// TODO: ! operator
	case "+":
		return arg1 + arg2
	case "*":
		return arg1 * arg2
	case "-":
		return arg2 - arg1
	case "^":
		return pow(arg2, arg1)
	case "/":
		return arg2 / arg1
	case "&&":
		return arg2 & arg1
	case "||":
		return arg2 | arg1
	case ">>":
		return arg2 >> arg1
	case "<<":
		return arg2 << arg1
	default:
		errors.New("bad operator")
	}
	return -1
}
