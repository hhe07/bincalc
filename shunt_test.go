package main

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	/* 
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
	sTests := map[string][]string{
		"2+3": []string{"2", "+", "3"},
		"2 + 3": []string{"2", "+", "3"},
		"4+32": []string{"4", "+", "32"},
		"6+4+2": []string{"6", "+", "4", "+", "2"},
		"2-3": []string{"2", "-", "3"},
		"2*3":[]string{"2", "*", "3"},
		"2/3":[]string{"2", "/", "3"},
		"2^3":[]string{"2", "^", "3"},
		"2&&3":[]string{"2", "&&", "3"},
		"2||3":[]string{"2", "||", "3"},
		"2<<3":[]string{"2", "<<", "3"},
		"2>>3":[]string{"2", ">>", "3"},
		"!2":[]string{"!", "2"},
		"(2)+((3)*6)": []string{"(","2",")","+","(","(","3",")","*","6",")"},
	}
	for in, out := range sTests{
		tkns, err := Tokenizer(in)
		if err != nil{
			t.Log("error should be nil", err)
			t.Fail()
		}
		for idx, elem := range tkns{
			if elem != out[idx]{
				t.Log("should be: ", out, "got: ", tkns)
				t.Fail()
			}
		}
	}
	// Tests that should fail:
	fTests := []string{
		"2^|3",
		"(2+3",
	}
	for _, ftest := range fTests{
		_, err := Tokenizer(ftest)
		if err == nil{
			t.Log("error should not be nil")
			t.Fail()
		}
	}

}



func TestCharType(t *testing.T){
	/*
		Testing cases:
		- Number: all acceptable/non
		- Function: acceptable/non
		- Operator: acceptable/non
		- Paren: acceptable/non
	*/
	sTests := map[TokenType][]rune{
		num: []rune{'0', '1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'},
		funct: []rune{'!'},
		operator: []rune{'^', '*', '/', '+', '-', '&', '|', '<', '>'},
		lparen: []rune{'('},
		rparen: []rune{')'},
	}
	fTests := []rune{
		'z',
		'~',
		'?',
	}
	for tType, cases := range sTests{
		for _, char := range cases{
			res := charType(char)
			if res != tType{
				t.Log("incorrect type:", char, "expected:", tType, "actual:",res)
				t.Fail()
			}
		}
	}
	for _, fCase := range fTests{
		res := charType(fCase)
		if res != -1{
			t.Log("incorrect type:", fCase, "expected:", -1, "actual:",res)
			t.Fail()
		}
	}
}

func TestIsNumber(t *testing.T) {
	/* 
		Testing cases:
		- bin: normal/abnormal (other numbers and a-f and other characters)
		- dec: normal/abnormal (a-f and other characters)
		- hex: normal/abnormal (other characters)
		- x
		- non-acceptable cases regardless
	*/
	tests := map[string]bool{
		"0b01": true,
		"0b02": false,
		"0b09": false,
		"0b0a": false,
		"0b0f": false,
		"0b0/": false,
		"0123456789": true,
		"0123456789a": false,
		"0123456789f": false,
		"0123456789/": false,
		"0x0123456789abcdef": true,
		"0x0123456789abcde/": false,
		"x": false,
		"/12345": false,
		"2gfdsg*": false,
	}
	for tc, er := range tests{
		res, err := isNumber(tc)
		if err != nil{
			t.Log(err)
			t.Fail()
		}
		if res != er{
			t.Log(tc, "expected: ", er, "got: ", res)
			t.Fail()
		}
	}
	_, err := isNumber("")
	if err == nil{
		t.Log("did not handle empty string")
		t.Fail()
	}

}

func TestPrecedence(t *testing.T) {
	/*
		Testing cases:
		- Each of the operators on top of empty list
		- Each of operators on top but with higher precedence in list
	*/
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
		[]string{"||", "<<"}, // todo: these need brane
		[]string{"||", ">>"},
		
	}

	fTests := [][]string{
		[]string{"^","^"},
		[]string{"*","^"},
		[]string{"+","/"},
		[]string{"&&","-"},
		[]string{"<<","^"},
		[]string{"<<","||"},
		[]string{">>","||"},
	}

	for _, test := range sTests{
		res := greatestPrecedence(test, test[1])
		if !res{
			t.Log("bad test: ", test, "expected: true, got: ", res)
			t.Fail()
		}
	}
	for _, test := range fTests{
		res := greatestPrecedence(test, test[1])
		if res{
			t.Log("bad test: ", test, "expected: false, got: ", res)
			t.Fail()
		}
	}
}


type OperatorTest struct{
	result int
	operator string
	numStack []int
}


func TestOperatorEval(t *testing.T) {
	/*
		Testing cases:
		- test each operator
		- test that incorrect operators are identified
	*/

	sTests := []OperatorTest{
		OperatorTest{
			result: 5,
			operator: "+",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 6,
			operator: "*",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: -1,
			operator: "-",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 8,
			operator: "^",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 1,
			operator: "/",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 2,
			operator: "&&",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 3,
			operator: "||",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 0,
			operator: ">>",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 16,
			operator: "<<",
			numStack: []int{2, 3},
		},
		OperatorTest{
			result: 1,
			operator: "!",
			numStack: []int{2},
		},
	}
	
	fTests := []string{
		"((",
		"JK",
		"~~",
	}
	
	for _, test := range sTests{
		res := operatorEval(test.numStack, test.operator)
		if res[0] != test.result{
			t.Log("bad test: ", test.operator, "expected: ", test.result, "got: ", res[0])
			t.Fail()
		}
	}
	for _, test := range fTests{
		res := operatorEval([]int{2,3}, test)
		if res != nil{
			t.Log("did not detect invalid operator ", test)
			t.Fail()
		}
	}

	
}