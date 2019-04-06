package main

import (
	"fmt"
	"strconv"
)

//TransMask is the bitmask for 0001 0110
const TransMask = 22

// ErrorCheck do the check
func ErrorCheck(err error) {
	if err != nil {
		fmt.Printf("Error while reading %v /n", err)
		panic("Stop")
	}
}

func count(input string) string {
	result := []rune{}
	for _, v := range input {
		result = append(result, v^TransMask)
	}
	return string(result)
}

func main() {
	var inputBuffer string
	fmt.Scanln(&inputBuffer)
	testTime, err := strconv.Atoi(inputBuffer)
	ErrorCheck(err)
	for i := 0; i < testTime; i++ {
		{
			var caseBuffer string
			fmt.Scanln(&caseBuffer)
			//useless -> drop
		}
		var caseBuffer string
		fmt.Scanln(&caseBuffer)
		ErrorCheck(err)
		result := count(caseBuffer)
		fmt.Printf("Case #%d: %s\n", i+1, result)
	}
}
