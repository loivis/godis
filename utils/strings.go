package utils

import "strings"
import "strconv"

// TrimAtoi ...
func TrimAtoi(s string) int {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Trim(s, " ")
	n, err := strconv.Atoi(s)
	CheckError(err)
	return n
}
