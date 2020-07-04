package raindrops

import "strconv"

//Convert method is used to check the factors of the given  number
func Convert(num int) (s string) {
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		s = strconv.Itoa(num)
		return s
	}
	if num%3 == 0 {
		s += "Pling"
	}
	if num%5 == 0 {
		s += "Plang"
	}

	if num%7 == 0 {
		s += "Plong"

	}

	return
}
