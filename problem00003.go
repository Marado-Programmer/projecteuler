package main

import "fmt"

func primes(n int) []int {
	p := []int{}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			n /= i
			p = append(p, i)
		}
	}
	return p
}

func lastElement(s []int) int {
	return s[len(s)-1]
}

func main() {
	fmt.Print(lastElement(primes(600851475143)))
}
