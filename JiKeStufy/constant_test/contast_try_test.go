package constant_test

import (
	"testing"
)
const(
	Monday = iota + 1
	Tuesday
	Wednesday
	Firday
)

const(
	Readable = 1 << iota
	Writable
	
)


func TestConstantTry(t *testing.T){
	t.Log(Monday,Tuesday)
}

func TestConstantTry1(t *testing.T){
	a :=7
	t.Log(a & Readable == Readable,a & Writable == Writable)
}