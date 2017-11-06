package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x;
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Printf("Hello, Denisa! Esti ok?\n")
	fmt.Printf(stringutil.Reverse("\n!oG ,olleH"))
	fmt.Println(add(43, 12))
	a, b := swap("Denisa", "Profir");
	fmt.Println(a,b);
	fmt.Println(split(17))
}
