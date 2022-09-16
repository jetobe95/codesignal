package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("Result %v\n", solution([]int{2, 4, 7}))

}

func solution(a []int) int {
	totals := []float64{}

	for _, v := range a {
		innerSum := 0.0
		for _, c := range a {
			abs := math.Abs(float64(c - v))
			innerSum += abs
		}
		totals = append(totals, float64(innerSum))

	}
	min := math.Inf(0)
	minIndex := 0
	for i, v := range totals {
		if v < min {
			min = v
			minIndex = i
		}

	}

	return int(a[minIndex])
}
