package main

func main() {
	// START OMIT
	for {
		// use this approach!
	}

	for ;; {
		// it works...
		// ... but don't do it!
		//     (and gofmt will rewrite it)
	}
	// END OMIT
}
