package main

import (
	"fmt"
)

func NumFibonacci(numFibonacci []uint64, l uint64) []uint64 {
	if l == 2 {
		return numFibonacci
	}
	numFibonacci = append(numFibonacci, numFibonacci[len(numFibonacci)-1]+numFibonacci[len(numFibonacci)-2])
	return NumFibonacci(numFibonacci, l-1)
}

func main() {
	var l uint64
	numFibonacci := []uint64{0, 1}

	fmt.Printf("\nВведите длину Числа Фибоначчи [0 1 ... l], где l = ")
	_, err := fmt.Scanln(&l)
	if err != nil || l <= uint64(len(numFibonacci)) {
		return
	}

	fmt.Println(NumFibonacci(numFibonacci, l))
}
