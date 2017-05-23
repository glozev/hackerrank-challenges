package main

import "fmt"

func SolveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = SolveMeFirst(a, b)
	fmt.Println(res)
}
