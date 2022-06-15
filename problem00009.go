package main

import (
	"fmt"
	"math"
)

var n, sum, prod int
var h float64

func foo(num int) {
	for c1 := 1; sum != num; c1++ {
		for c2 := c1+1; sum < num; c2++ {
			h = math.Sqrt(float64(c1*c1 + c2*c2))
			sum = c1 + c2 + int(h)
			if math.Mod(h, 1) != 0 {
				continue
			} else {
				h := int(h)
				sum = c1 + c2 + h
				prod = c1 * c2 * h
				if sum == num {
					fmt.Print(c1, ",", c2, ",", h, ";\t")
					fmt.Println("sum", sum, "; prod", prod)
				}
			}
		}
		sum = 0
	}
}

func main() {
	n = 1000
	foo(n)
}

