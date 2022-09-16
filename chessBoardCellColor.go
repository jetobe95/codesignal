package main

import (
	"strconv"
	"strings"
)

// func solution(cel1 string, cel2 string) bool {
// 	if cel1 == cel2 {
// 		return true
// 	}
// 	c1Letter, c1Num := splitValues(cel1)
// 	c2Letter, c2Num := splitValues(cel2)

// 	return hasColor(c1Letter, c1Num) && hasColor(c2Letter, c2Num)
// }

func splitValues(cel string) (letterPos, num int) {
	cel1Slice := strings.Split(cel, ``)
	num, _ = strconv.Atoi(cel1Slice[1])
	letterPos = int([]byte(cel1Slice[0])[0])
	return
}

func hasColor(x, y int) bool {
	isEven1, IsEven2 := x%2 == 0, y%2 == 0
	return !(isEven1 || IsEven2)
}
