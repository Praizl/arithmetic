package main

import (
	"fmt"

	"github.com/Praizl/arithmetic/factorial"
	"github.com/Praizl/arithmetic/prime"
)

func main() {
	fmt.Println(prime.IsPrime(13))
	fmt.Println(factorial.Factorial(69))
}
