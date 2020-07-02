package leap

//IsLeapYear is function used to check whether the given year is leap year or not
func IsLeapYear(x int) bool {

	if (x%100 != 0 && x%4 == 0) || (x%400 == 0) {

		return true
	}

	return false
}
