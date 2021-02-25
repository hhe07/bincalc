package bincalc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Tokenizer() {}

// or have tokens separated by spaces when entering args?

/*
type Token struct {
	Text          string
	Precedence    int
	Associativity bool // false: left, true: right
}

// Will the functions inherently be left associative?
// Will the functions inherently be lowest precedence

func TokenType(token string) (Token, error) {
	// Operators: + - * / ^ && || << >>
	// Functions: !
	operators := [10]Token{
		Token{Text: "^", Precedence: 4, Associativity: true},
		Token{Text: "*", Precedence: 3, Associativity: false},
		Token{Text: "/", Precedence: 3, Associativity: false},
		Token{Text: "+", Precedence: 2, Associativity: false},
		Token{Text: "-", Precedence: 2, Associativity: false},
		Token{Text: "&&", Precedence: 0, Associativity: false},
		Token{Text: "||", Precedence: 0, Associativity: false},
		Token{Text: "<<", Precedence: 0, Associativity: true},
		Token{Text: ">>", Precedence: 0, Associativity: true},
		//		Token{Text: "!", Precedence: 1, Associativity: false},
	}

	// Handle operator or function
	for _, t := range operators {
		if t.Text == token {
			return t, nil
		}
	}
	return Token{}, errors.New("Token out of range")
}
*/
func inRangeNInc(low, high, test int) bool {
	// Non-inclusive inRange
	return (low < test) && (test < high)
}

func inRangeInc(low, high, test int) bool {
	// Inclusive inRange
	return (low <= test) && (test <= high)
}

func isNumber(token string) (bool, error) {
	// Takes only lowercased strings
	if len(token) <= 0 {
		return false, errors.New("No length")
	}

	for _, cp := range token {
		// CP is the unicode codepoint
		// If not a number                  not lowercase                      not x (possible for 0x)
		if !(inRangeInc(48, 57, int(cp)) || (inRangeInc(97, 102, int(cp)) || int(cp) == 120)) {
			return false, nil
		}
	}
	return true, nil

}

func greatestPrecedence(opStack []string) bool {
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
	rightAs := [3]string{"^", "<<", ">>"}
	topPrec := opPrec[opStack[0]]
	for _, op := range opStack {
		currPrec := opPrec[op]

		isRightAs := false
		for _, r := range rightAs {
			if r == op {
				isRightAs = true
				break
			}
		}

		if !((currPrec < topPrec) || ((currPrec == topPrec) && !isRightAs)) {
			return false
		}
	}
	return true
}

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

		// case: number
		if isNum {
			ret = append([]string{token}, ret...)
		} else if token == "!" {
			operatorStack = append([]string{token}, operatorStack...)
		} else if token != "(" && token != ")" {
			for len(operatorStack) >= 0 && (greatestPrecedence(operatorStack) && operatorStack[0] != "(") {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret, operator)
			}
			operatorStack = append([]string{token}, operatorStack...)
		} else if token == "(" {
			operatorStack = append([]string{token}, operatorStack...)
		} else if token == ")" {
			for operatorStack[0] != "(" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret, operator)
				if (operatorStack[0] != "(") && (len(operatorStack) == 1) {
					return nil, errors.New("Mismatched parenthesis")
				}
			}
			// todo: if runs out w/o finding left paren, then mismatched
			if operatorStack[0] == "(" {
				operatorStack = operatorStack[1:]
			}
			if operatorStack[0] == "!" {
				operator, operatorStack = operatorStack[0], operatorStack[1:]
				ret = append(ret, operator)
			}

		}
	}
	if len(operatorStack) > 0 {
		if (operatorStack[0] == "(") || (operatorStack[0] == ")") {
			return nil, errors.New("Mismatched parenthesis")
		}
		for len(operatorStack) > 0 {
			// todo: if top is parenthesis, mismatched
			operator, operatorStack = operatorStack[0], operatorStack[1:]
			ret = append(ret, operator)
		}
	}

	return ret, nil
}

func pow(base, pow int) int {
	ret := base
	for x := 1; x < pow; x++ {
		ret *= base
	}
	return ret
}

func RPNInterpret(tokens []string) (int, error) {
	evalStack := make([]int, 0)
	var arg1 int
	var arg2 int
	for len(tokens) > 0 {
		top, tokens := tokens[0], tokens[1:]
		isNum, err := isNumber(top)
		if err != nil {
			errors.New("Bad token")
		}
		if isNum {
			iConv, err := strconv.Atoi(top)
			if err != nil {
				errors.New("Bad int conversion")
			}
			evalStack = append([]int{iConv}, evalStack...)
		} else {
			switch top {
			case "+":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg1 + arg2}, evalStack...)
			case "*":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg1 * arg2}, evalStack...)
			case "-":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg2 - arg1}, evalStack...)
			case "^":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{pow(arg2, arg1)}, evalStack...)
			case "/":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				if arg1 == 0 {
					return -1, errors.New("Cannot div by 0")
				}
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg2 / arg1}, evalStack...)
			case "&&":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg2 & arg1}, evalStack...)
			case "||":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg2 | arg1}, evalStack...)
			case ">>":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg2 >> arg1}, evalStack...)
			case "<<":
				if len(evalStack) < 2 {
					return -1, errors.New("Bad expr")
				}
				arg1, evalStack = evalStack[0], evalStack[1:]
				arg2, evalStack = evalStack[0], evalStack[1:]
				evalStack = append([]int{arg2 << arg1}, evalStack...)
			default:
				return -1, errors.New("unknown command")
			}
		}

	}

	if len(evalStack) == 1 {
		return evalStack[0], nil
	}
	return -1, errors.New("Something went wrong")
}
