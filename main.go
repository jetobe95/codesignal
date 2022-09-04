package main

import "fmt"

func main() {

	// image := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	// image := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	// image := [][]int{{1, 1, 1}, {1, 7, 1}, {1, 1, 1}}
	// image := [][]int{{7, 4, 0, 1},
	// 	{5, 6, 2, 2},
	// 	{6, 10, 7, 8},
	// 	{1, 4, 2, 0}}
	image := [][]int{{36, 0, 18, 9},
		{27, 54, 9, 0},
		{81, 63, 72, 45}}
	imgOut := solution(image)

	fmt.Println(imgOut)

}

func solution(image [][]int) [][]int {
	imgOut := [][]int{}
	totalGroups := len(image[0])
	rowsCount := len(image)
	colsCount := len(image[0])

	fmt.Println(rowsCount, colsCount)

	for i := 1; i < totalGroups-1; i++ {
		rowAvg := []int{}
		for i2 := 1; i2 < totalGroups-1; i2++ {
			sum := 0
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					sum += image[i-1+j][i2-1+k]
				}
			}
			sum /= 9

			rowAvg = append(rowAvg, sum)
		}
		imgOut = append(imgOut, rowAvg)

	}

	return imgOut
}
