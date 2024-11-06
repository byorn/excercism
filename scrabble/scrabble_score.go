package scrabble

func Score(word string) int {
	var m = make(map[string]int)
	m["AEIOULNRST"] = 1
	m["DG"] = 2
	m["BCMP"] = 3
	m["FHVWY"] = 4
	m["K"] = 5
	m["JX"] = 8
	m["QZ"] = 10

	panic("Please implement the Score function")
}
