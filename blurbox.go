package main

func blurBox(image [][]int) [][]int {
	// 1. start calc 3x3 center indices i.e for 5x5 image [1,1]	[1,2] [1,3] [2,1][2,2][2,3]
	// 2. from center indices calc borders
	// 3. calc avg
	imgOut := [][]int{}
	colsCount := len(image[0])
	rowsCount := len(image)
	for i := 1; i < rowsCount-1; i++ {
		rowAvg := []int{}
		for i2 := 1; i2 < colsCount-1; i2++ {
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
