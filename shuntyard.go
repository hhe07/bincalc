package main

import (
	"errors"
	"fmt"
	"strings"
)

func Tokenizer(input string) ([]string, error) {
	ret := make([]string, 0)
	token := ""
	prevType := 0 // 0: num, 1: func, 2: operator, 3/4: l/r paren
	// todo: check parenthesis count
	// todo: check for mistyped operations, esp. boolean ones
	for idx, cp := range input {
		if idx == 0 {
			prevType = charType(cp)
		}
		if prevType != charType(cp) {
			ret = append(ret, token)
			prevType = charType(cp)
			token = ""
		}

		token += string(cp)

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
	/*
		fmt.Println(op, opStack[0])
		fmt.Println(currPrec, topPrec)
		fmt.Println(currPrec < topPrec)
	*/
	return (currPrec > topPrec) || ((currPrec == topPrec) && !isRightAs)
}

// TODO: How exactly do I make a shuntyard integrated w/ evaluator? (mostly in operation step)
// Answer: instead of moving operators to ret stack, evaluate the operator
func ShuntYard(tokens []string) ([]string, error) {
	ret := make([]string, 0)
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
		fmt.Println("pass no", i)
		fmt.Println(token)
		//fmt.Println(ret, operatorStack)
		//fmt.Println()
		// case: number
		if isNum {
			fmt.Println("n")
			ret = append(ret, token)
			// case: function
		} else if token == "!" {
			fmt.Println("f")
			operatorStack = append([]string{token}, operatorStack...)
			// case: operator
		} else if token != "(" && token != ")" {
			fmt.Println("o")
			for len(operatorStack) > 1 && (greatestPrecedence(operatorStack, token) && operatorStack[0] != "(") {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret, operator)
			}
			operatorStack = append([]string{token}, operatorStack...)
			// case: lparen
		} else if token == "(" {
			fmt.Println("lp")
			operatorStack = append([]string{token}, operatorStack...)
			// case: rparen
		} else if token == ")" {
			fmt.Println("rp")
			for operatorStack[0] != "(" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret, operator)
				if (operatorStack[0] != "(") && (len(operatorStack) == 1) {
					return nil, errors.New("Mismatched parenthesis")
				}
			}
			if operatorStack[0] == "(" {
				operatorStack = operatorStack[1:]
			}
			if operatorStack[0] == "!" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append([]string{operator}, ret...)
			}

		}
	}
	if len(operatorStack) > 0 {
		if (operatorStack[0] == "(") || (operatorStack[0] == ")") {
			return nil, errors.New("Mismatched parenthesis")
		}
		for len(operatorStack) > 0 {
			operator, operatorStack = operatorStack[0], operatorStack[1:]
			ret = append(ret, operator)
		}
	}

	return ret, nil
}

func main() {
	/*
		x := greatestPrecedence([]string{"*", "+"})
		fmt.Println(x)
	*/
	x, err := ShuntYard([]string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2", "^", "3"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

	/*
		x, err := Tokenizer("32 + 4 * 2 / (1 - 5) ^ 2 ^ 3")
		//x, err := Tokenizer("3+4")
		//x, err := Tokenizer("32+4*2/(1-5)^2^3")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(x)
		fmt.Println(len(x))
	*/
}
