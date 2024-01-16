package funcPkg

import (
	"fmt"
	"math/rand"
)

type transformFn func(int) int

// passing func as parameter
func transformNumbers(arr *[]int, transform transformFn) []int {
	resArr := make([]int, len(*arr))
	for i, val := range *arr {
		resArr[i] = transform(val)
	}
	return resArr
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}

// Returning func as return value
func getRandomTransformFunc() transformFn {
	if getRandomValue() <= 5 {
		fmt.Println("Double is selected")
		return double
	} else {
		fmt.Println("Triple is selected")
		return triple
	}
}

func getRandomValue() int {
	return rand.Intn(10)
}

// closures
func getTransformFuncByFactor(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}

// variadic functions
func sum(initialValue int, numbers ...int) int {
	sum := initialValue
	for _, val := range numbers {
		sum += val
	}
	return sum
}
