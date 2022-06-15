package main

import (
	"fmt"
)

var (
	n int
)

func sumOfPrimesBelow(natural int) (int) {
	var primes []int
	for i := 2; i < natural; i++ {
		if primes == nil {
			primes = append(primes, i)
		} else {
			var isPrime bool = true
			for _, j := range primes {
				if i % j == 0 {
					isPrime = false
				}
			}
			if isPrime {
				primes = append(primes, i)
			}
		}
	}
	var sum int
	fmt.Println(primes)
	for _, i := range primes {
		sum += i
	}
	return sum
}

func main() {
	n = 2000000
	fmt.Print(sumOfPrimesBelow(n))
}
