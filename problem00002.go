package main

import "fmt"

func main() {
	sum, n1, n2, max := 0, 1, 2, int(4e+6)
	for n2 < max {
		if n2%2 == 0 {
			sum += n2
		}
		n2 += n1
		n1 = n2 - n1
	}
	fmt.Print(sum)
}
