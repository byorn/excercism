package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	strName := "you"

	if name != "" {
		strName = name
	}
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return "one for " + strName + " one for me."
}
