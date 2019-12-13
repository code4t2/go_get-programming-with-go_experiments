package main

import (
	"fmt"
	"math"
)

func main() {
	var integer int = math.MaxInt64 // default type is 'int', 64-bits on amd64, 32 bits on i86-32
	var integer8 int8 = math.MaxInt8
	var integer16 int16 = math.MinInt16
	var integer32 int32 = math.MinInt32
	var integer64 int64 = math.MaxInt64
	fmt.Println(integer, integer8, integer16, integer32, integer64)
	var uinteger uint = math.MaxUint64
	var uinteger8 uint8 = math.MaxUint8
	var uinteger16 uint16 = math.MaxUint16
	var uinteger32 uint32 = math.MaxUint32
	var uinteger64 uint64 = math.MaxUint64
	fmt.Println(uinteger, uinteger8, uinteger16, uinteger32, uinteger64)
}
