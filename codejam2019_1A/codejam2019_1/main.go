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
func main() {
	var inputBuffer string
	fmt.Scanln(&inputBuffer)
	testTime, err := strconv.Atoi(inputBuffer)
	ErrorCheck(err)
	for i := 0; i < testTime; i++ {
		var r, c int
		{
			fmt.Scan(&r)
		}

		{
			fmt.Scan(&c)
		}

		possibility, steps := count(r, c)
		if possibility == true {
			fmt.Printf("Case #%d: POSSIBLE\n", i+1)
			//fmt.Printf("Start verify %d, %d\n", r, c)
			verify(steps, r, c)
			if len(steps) != r*c {
				fmt.Printf("for case %d %d \n", r, c)
				panic("exit")
			}
			for _, step := range steps {
				fmt.Printf("%d %d\n", step[0], step[1])
			}
		} else {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
			//fmt.Printf("for case %d %d \n", r, c)
		}
	}
}

func verify(input [][]int, r int, c int) {
	for i := 0; i < len(input)-1; i++ {
		if input[i][0]-input[i][1] == input[i+1][0]-input[i+1][1] {
			fmt.Printf("%d, %d break with %d, %d\n",
				input[i][0],
				input[i][1],
				input[i+1][0],
				input[i+1][1])
			err := string("Case " + string(r) + " " + string(c) + " exit")
			panic(err)
		}
		if input[i][0]+input[i][1] == input[i+1][0]+input[i+1][1] {
			fmt.Printf("%d, %d break with %d, %d\n",
				input[i][0],
				input[i][1],
				input[i+1][0],
				input[i+1][1])
			err := string("Case " + string(r) + " " + string(c) + " exit")
			panic(err)
		}
		if input[i][0] == input[i+1][0] {
			fmt.Printf("%d, %d break with %d, %d\n",
				input[i][0],
				input[i][1],
				input[i+1][0],
				input[i+1][1])
			err := string("Case " + string(r) + " " + string(c) + " exit")
			panic(err)
		}
		if input[i][1] == input[i+1][1] {
			fmt.Printf("%d, %d break with %d, %d\n",
				input[i][0],
				input[i][1],
				input[i+1][0],
				input[i+1][1])
			err := string("Case " + string(r) + " " + string(c) + " exit")
			panic(err)
		}
	}
	return
}

func count(r int, c int) (bool, [][]int) {
	//fmt.Printf("%d %d \n", r, c)
	var result [][]int
	if r <= 3 && c <= 3 {
		//fmt.Printf("fail %d %d \n", r, c)
		return false, [][]int{}
	} else if r < 3 && c < 5 {
		return false, [][]int{}
	} else if r < 5 && c < 3 {
		return false, [][]int{}
	} else if r == 3 {
		return true, count3(r, c)
	} else if c == 3 {
		for i := 1; i <= r; i++ {
			result = append(result, []int{i, 1})
			result = append(result, []int{i + 2, 2})
			if result[len(result)-1][0] > r {
				result[len(result)-1][0] -= r
			}
			result = append(result, []int{i, 3})
		}
		return true, result
	}

	if r == 4 && c == 4 {
		result := [][]int{
			{1, 4}, {2, 2}, {3, 4}, {4, 2},
			{1, 1}, {2, 3}, {3, 1}, {4, 3},
			{1, 2}, {2, 4}, {3, 2}, {4, 4},
			{1, 3}, {2, 1}, {3, 3}, {4, 1}}
		return true, result
	}

	if r%2 == 0 && c%2 == 0 {
		if r < c {
			for i := 0; i < r; i += 2 {
				result = append(result, count2(r, c, i+1)...)
			}
		} else {
			for i := 0; i < c; i += 2 {
				result = append(result, revertXY(count2(c, r, i+1))...)
			}
		}
	} else if r%2 == 0 {
		for i := 0; i < r; i += 2 {
			result = append(result, count2(r, c, i+1)...)
		}
	} else if c%2 == 0 {
		for i := 0; i < c; i += 2 {
			result = append(result, revertXY(count2(c, r, i+1))...)
		}
	} else {
		result = append(result, count3(r, c)...)
		for i := 0; i < r-3; i += 2 {
			result = append(result, count2(r, c, i+4)...)
		}
	}
	return true, result
}

func count3(r int, c int) [][]int {
	var result [][]int
	for i := 1; i <= c; i++ {
		result = append(result, []int{1, i})
		result = append(result, []int{2, (i + 2)})
		if result[len(result)-1][1] > c {
			result[len(result)-1][1] -= c
		}
		result = append(result, []int{3, i})
	}
	return result
}

func count2(r int, c int, startR int) [][]int {
	var result [][]int
	for i := 1; i <= c; i++ {
		result = append(result, []int{startR, i + 2})
		if result[len(result)-1][1] > c {
			result[len(result)-1][1] -= c
		}
		result = append(result, []int{startR + 1, i})
	}
	return result
}

func revertXY(input [][]int) [][]int {
	for _, v := range input {
		v[0], v[1] = v[1], v[0]
	}
	return input
}
