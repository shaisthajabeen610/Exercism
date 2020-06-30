package hamming

import "errors"

//Distance is a function used to determine the hamming distance between two strings
func Distance(a, b string) (int, error) {
	var count int = 0
	var err1 error = nil

	if len(a) != len(b) {
		return 0, errors.New("unequal strings")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, err1

}
