package main

import "fmt"

func main() {
	var num int
	for i, cnt := 1, 0; cnt <= 10001; i++ {
		for j := 1; j <= i; j++ {
			if i % j == 0 && j != 1 && j != i {
				break
			}
			if j == i {
				cnt++
				num = i
			}
		}
	}
	fmt.Print(num)
}

