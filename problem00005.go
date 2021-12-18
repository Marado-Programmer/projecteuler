package main

import "fmt"

func main() {
	num, from, to, notFound := 0, 1, 20, true
	for val := 1; notFound; val++ {
		for i := to; i >= from; i-- {
			if !(val % i == 0) {
				break
			}
			fmt.Println(val, i)
			if i == from {
				fmt.Println("POG")
				num = val
				notFound = false
			}
		}
	}
	fmt.Print(num)
}

