package main

import (
	"fmt"
)

func main() {
	const (
		A = iota
		B
		C
		D = "HAHA"
		E
		F = 100
		G
		H = iota
		I
	)
	const (
		J = iota
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(G)
	fmt.Println(J)

	/*
		0
		1
		2
		HAHA
		HAHA
		100
		100
		7
		8
		100
		0
	*/
}
