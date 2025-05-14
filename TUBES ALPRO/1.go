package main

import "fmt"

func main() {
	var A, B, C int
	var D, P *int
	A = 3
	B = 4
	C = A + B
	D = &C
	fmt.Println(*D)
	P = &A
	*P = *D + *P
	fmt.Println(P)
	fmt.Println(A)

}
