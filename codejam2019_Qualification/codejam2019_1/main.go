package main

import (
	"fmt"
	"strconv"
)

// ErrorCheck do the check
func ErrorCheck(err error) {
	if err != nil {
		fmt.Printf("Error while reading %v /n", err)
		panic("Stop")
	}
}

func count(input int) (int, int) {
	//fmt.Printf("%v \n", input)
	left := []rune(strconv.Itoa(input - 1))
	right := []rune{'1'}
	//fmt.Printf("%v \n", left)
	//fmt.Printf("%v \n", left)
	if left[len(left)-1] == '4' {
		right[0]++
		left[len(left)-1]--
	}
	for i := len(left) - 2; i >= 0; i-- {
		right = append([]rune{'0'}, right...)
		if left[i] == '4' {
			left[i]--
			right[0]++
		}
		//fmt.Printf("%v /n", left)
	}
	resultLeft, _ := strconv.Atoi(string(left))
	resultRight, _ := strconv.Atoi(string(right))

	return resultLeft, resultRight
}

func main() {
	var inputBuffer string
	fmt.Scanln(&inputBuffer)
	testTime, err := strconv.Atoi(inputBuffer)
	ErrorCheck(err)
	for i := 0; i < testTime; i++ {
		var caseBuffer string
		fmt.Scanln(&caseBuffer)
		testCase, err := strconv.Atoi(caseBuffer)
		ErrorCheck(err)
		left, right := count(testCase)
		fmt.Printf("Case #%d: %d %d\n", i+1, left, right)
	}
}
