package main

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	/*
		Unit tests for Tokenizer
		Testing cases:
		- single-length num
		- multi-length num
		- single-length more than 2 nums
		- all the operators
		- two-part operators wrong
		- right paren matching
		- wrong paren matching
	*/
	// Tests that should succeed:
	// test if each string
	// tokenizes to the right list
	sTests := map[string][]string{
		"2+3":         []string{"2", "+", "3"},
		"2 + 3":       []string{"2", "+", "3"},
		"4+32":        []string{"4", "+", "32"},
		"6+4+2":       []string{"6", "+", "4", "+", "2"},
		"2-3":         []string{"2", "-", "3"},
		"2*3":         []string{"2", "*", "3"},
		"2/3":         []string{"2", "/", "3"},
		"2^3":         []string{"2", "^", "3"},
		"2&&3":        []string{"2", "&&", "3"},
		"2||3":        []string{"2", "||", "3"},
		"2<<3":        []string{"2", "<<", "3"},
		"2>>3":        []string{"2", ">>", "3"},
		"!2":          []string{"!", "2"},
		"(2)+((3)*6)": []string{"(", "2", ")", "+", "(", "(", "3", ")", "*", "6", ")"},
	}
	// Tests that should fail:
	// these should not be accepted
	// because of input errors
	fTests := []string{
		"2^|3",
		"(2+3",
	}
	// run successful tests
	for in, out := range sTests {
		// in is input that should be tokenized, out is expeced result
		// get tokenizer result
		tkns, err := Tokenizer(in)
		if err != nil {
			t.Log("error should be nil", err)
			t.Fail()
		}
		// compare elements of out and tkns
		for idx, elem := range tkns {
			if elem != out[idx] {
				t.Log("should be: ", out, "got: ", tkns)
				t.Fail()
			}
		}
	}

	// run failing tests
	for _, ftest := range fTests {
		// idx not necessary
		// take tokens
		_, err := Tokenizer(ftest)
		// should have an error, so err should not be nil.
		if err == nil {
			t.Log("error should not be nil")
			t.Fail()
		}
	}

}

func TestCharType(t *testing.T) {
	/*
		Unit tests for charType
		Testing cases:
		- Number: all acceptable/non
		- Function: acceptable/non
		- Operator: acceptable/non
		- Paren: acceptable/non
	*/
	// Tests that should succeed:
	// test if each list of characters
	// falls under the indicated category
	sTests := map[TokenType][]rune{
		num:      []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'},
		funct:    []rune{'!'},
		operator: []rune{'^', '*', '/', '+', '-', '&', '|', '<', '>'},
		lparen:   []rune{'('},
		rparen:   []rune{')'},
	}
	// Tests that should fail:
	// these characters should not be accepted
	fTests := []rune{
		'z',
		'~',
		'?',
	}
	// run successful tests
	for tType, cases := range sTests {
		// tType is the expected result, cases are the test cases
		// this is because a slice cannot be the key to a map.
		for _, char := range cases {
			// char is the actual case
			res := charType(char)
			// compare result of charType to expected result
			if res != tType {
				t.Log("incorrect type:", char, "expected:", tType, "actual:", res)
				t.Fail()
			}
		}
	}
	// run failing tests
	for _, fCase := range fTests {
		// idx not necessary
		// take chartype
		res := charType(fCase)
		// if it isn't invalid chartype, fail.
		if res != -1 {
			t.Log("incorrect type:", fCase, "expected:", -1, "actual:", res)
			t.Fail()
		}
	}
}

func TestIsNumber(t *testing.T) {
	/*
		Unit tests for isNumber
		Testing cases:
		- bin: normal/abnormal (other numbers and a-f and other characters)
		- dec: normal/abnormal (a-f and other characters)
		- hex: normal/abnormal (other characters)
		- x
		- non-acceptable cases regardless
	*/
	// successful and failed tests are stored as one
	// second part of map is expected result of isNumber on first
	tests := map[string]bool{
		"0b01":               true,
		"0b02":               false,
		"0b09":               false,
		"0b0a":               false,
		"0b0f":               false,
		"0b0/":               false,
		"0123456789":         true,
		"0123456789a":        false,
		"0123456789f":        false,
		"0123456789/":        false,
		"0x0123456789abcdef": true,
		"0x0123456789abcde/": false,
		"x":                  false,
		"/12345":             false,
		"2gfdsg*":            false,
	}
	// iterate over and evaluate tests
	for tc, er := range tests {
		// tc is test case, er is expected result
		// get whether tc is number
		res, err := isNumber(tc)
		// no errors should be caused, fail in that case.
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		// check result versus expected
		if res != er {
			t.Log(tc, "expected: ", er, "got: ", res)
			t.Fail()
		}
	}
	// check for empty string: should actually cause error
	_, err := isNumber("")
	// fail if no error
	if err == nil {
		t.Log("did not handle empty string")
		t.Fail()
	}

}

func TestPrecedence(t *testing.T) {
	/*
		Unit tests for greatestPrecedence
		Testing cases:
		- Each of the operators in higher precedence situation
		- Each of operators in lower precedence situation
	*/
	// Tests that should succeed:
	// operator at top of stack should be
	// "greater" than second element
	sTests := [][]string{
		[]string{"^", "*"},
		[]string{"*", "*"},
		[]string{"*", "/"},
		[]string{"/", "+"},
		[]string{"+", "+"},
		[]string{"+", "-"},
		[]string{"-", "&&"},
		[]string{"&&", "&&"},
		[]string{"&&", "||"},
		[]string{"<<", "||"},
		[]string{"<<", "||"},
	}

	// Tests that should fail:
	// operator at top of stack should not be
	// "greater" than second element
	fTests := [][]string{
		[]string{"^", "^"},
		[]string{"*", "^"},
		[]string{"+", "/"},
		[]string{"&&", "-"},
		[]string{"<<", "^"},
		[]string{"||", "<<"},
		[]string{"||", ">>"},
	}

	// run successful tests
	for _, test := range sTests {
		// idx not necessary
		// take precedence check
		res := greatestPrecedence(test, test[1])
		// expected result is true, so if false, fail.
		if !res {
			t.Log("bad test: ", test, "expected: true, got: ", res)
			t.Fail()
		}
	}
	// run failing tests
	for _, test := range fTests {
		// take precedence check
		res := greatestPrecedence(test, test[1])
		// expected result is false, so if true, fail.
		if res {
			t.Log("bad test: ", test, "expected: false, got: ", res)
			t.Fail()
		}
	}
}

// small data structure to define test cases
// for operatorEval
type OperatorTest struct {
	result   int
	operator string
	numStack []int
}

func TestOperatorEval(t *testing.T) {
	/*
		Unit tests for operatorEval
		Testing cases:
		- test each operator
		- test that incorrect operators are identified
	*/

	// Tests that should succeed:
	// result is the expected result
	// of applying the operator
	// to the numStack
	sTests := []OperatorTest{
		OperatorTest{
			result:   5,
			operator: "+",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   6,
			operator: "*",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   -1,
			operator: "-",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   8,
			operator: "^",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   1,
			operator: "/",
			numStack: []int{3, 2},
		},
		OperatorTest{
			result:   2,
			operator: "&&",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   3,
			operator: "||",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   0,
			operator: ">>",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   16,
			operator: "<<",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result:   1,
			operator: "!",
			numStack: []int{2},
		},
	}

	// Tests that should fail:
	// operators that aren't defined
	fTests := []string{
		"((",
		"JK",
		"~~",
	}

	// run successful tests
	for _, test := range sTests {
		// get evaluation
		res := operatorEval(test.numStack, test.operator)
		// if the result is contrary to expected, fail
		if res[0] != test.result {
			t.Log("bad test: ", test.operator, "expected: ", test.result, "got: ", res[0])
			t.Fail()
		}
	}
	// run failing tests
	for _, test := range fTests {
		// get evaluation
		res := operatorEval([]int{2, 3}, test)
		// if result isn't nil, bad operator wasn't detected, so fail.
		if res != nil {
			t.Log("did not detect invalid operator ", test)
			t.Fail()
		}
	}

}
