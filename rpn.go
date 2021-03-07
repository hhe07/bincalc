package main

import (
	"errors"
	"strconv"
)

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
	var top string
	for len(tokens) > 0 {
		top, tokens = tokens[0], tokens[1:]
		isNum, err := isNumber(top)
		if err != nil {
			return -1, errors.New("Bad token")
		}
		if isNum {
			iConv, err := strconv.Atoi(top)
			if err != nil {
				return -1, errors.New("Bad int conversion")
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
