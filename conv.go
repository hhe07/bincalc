package bincalc

import (
	"strconv"
	"strings"
)

// Number: Type for representing one of the vals
type Number struct {
	bin int64
	dec int64
	hex int64
}

func newNum(numString string) *Number {
	var ret *Number
	if strings.HasPrefix(numString, "0x") {
		decHex := strconv.ParseInt(numString, 0, 64)
		ret = &Number{hex: decHex}
	}
	if strings.HasPrefix(numString, "0b") {
		decBin := strconv.ParseInt(numString, 0, 64)
		ret = &Number{hex: decHex}
	}
	if strings.HasPrefix(numString, "0d") {
		decDec := strconv.ParseInt(numString, 0, 64)
		ret = &Number{hex: decHex}
	}
	return ret
}

func DecToBX(input, base int) string {
	symbols := [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	ret := ""
	quot := input
	for quot > 0 {
		val := quot
		ret += string(quot % base)
		quot = (quot - (quot % base)) / base
	}
	return ret
}
