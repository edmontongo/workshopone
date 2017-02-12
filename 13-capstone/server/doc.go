// Command server provides...
package main

import "io"

func writeDocumentation(wr io.Writer) {
	wr.Write([]byte("Documentation!"))
}
