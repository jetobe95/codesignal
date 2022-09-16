package main

func ArrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	outputArray := []int{}
	for _, value := range inputArray {
		if elemToReplace == value {
			outputArray = append(outputArray, substitutionElem)
		} else {
			outputArray = append(outputArray, value)

		}
	}
	return outputArray
}
