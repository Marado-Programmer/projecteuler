package main

import "fmt"

func isPalindromic(n int) bool {
	var splitedNum []int
	
	for ;n != 0; n /= 10 {
		splitedNum = append(splitedNum, n % 10)
	}

	_len := len(splitedNum)

	for i := 1; i < _len; i++ {
		if !(splitedNum[_len-i] == splitedNum[i-1]) {
			return false
		}
		if _len-i == i-1 || _len/2 == _len/2-1{
			break
		}
	}

	return true
}

func main() {
	var val int
	i := 999
	for ; i > 99; i-- {
		for j := 999 ; j > 99; j-- {
			if isPalindromic(i*j) && i*j > val {
				val = i*j
			}
		}
	}
	fmt.Println(i, "*", val/i, "=", val)
}
