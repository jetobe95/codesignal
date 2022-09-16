package main

func alphabeticShift(inputString string) string {
	byteArray := []byte(inputString)
	output := ""
	for _, value := range byteArray {
		tmp := value + 1

		if charV := int(value); charV == 122 {
			tmp = byte('a')
		}
		output += string(tmp)
	}

	return output
}
