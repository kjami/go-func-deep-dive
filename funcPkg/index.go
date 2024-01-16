package funcPkg

import (
	"fmt"
	"math"
)

func Practice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original: ", slice)
	fmt.Println("Doubled: ", transformNumbers(&slice, double))
	fmt.Println("Tripled: ", transformNumbers(&slice, triple))
	fmt.Println("Random: ", transformNumbers(&slice, getRandomTransformFunc()))
	// Passing anonymous functions
	fmt.Println("Power of 2: ", transformNumbers(&slice, func(number int) int {
		return int(math.Pow(float64(number), 2.0))
	}))
	// Usage of closures
	fmt.Println("Multiple by 50: ", transformNumbers(&slice, getTransformFuncByFactor(50)))
	// Calling variadic function
	fmt.Println("Sum: ", sum(0, slice...))
	fmt.Println("Sum of a random slice: ", sum(1, 2, 3, 4, 5))
	// Factorial
	fmt.Println("Factorial of 6: ", factorial(6))
}
