package twofer

func ShareWith(name string) string {

	x := ""
	if name == "" {
		x = "you"
	} else {
		x = name
	}
	return "One for " + x + ", one for me."
}
