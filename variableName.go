package main

import "regexp"

func variableName(name string) bool {
	/// Correct variable names consist only of English letters, digits and underscores
	/// and they can't start with a digit.
	validRegex, _ := regexp.Compile(`^[^\d\s][A-Za-z\d_]*$`)
	return validRegex.MatchString(name)
}
