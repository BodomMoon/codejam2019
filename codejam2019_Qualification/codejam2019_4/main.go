package main

import (
	"fmt"
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
	
}

func main() {
	var testTime int
	_, err := fmt.Scanf("%d", &testTime)
	ErrorCheck(err)
	for i := 0; i < testTime; i++ {
		var n, b, f int
		_, err := fmt.Scanf("%d %d %d", &n, &b, &f)
		ErrorCheck(err)

		result := count(caseBuffer)
		fmt.Printf("Case #%d: %s\n", i+1, result)
	}
}
