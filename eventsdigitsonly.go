package main

import (
	"strconv"
	"strings"
)

func eventDigitsOnly(n int) bool {
	for _, value := range strings.Split(strconv.Itoa(n), "") {
		if in, _ := strconv.Atoi(value); in%2 != 0 {
			return false
		}

	}
	return true
}
