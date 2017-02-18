// Command server provides a basic RESTesque API for participants to practice
// writing simple Go programs against.
//
// Basic HTTP API sketch
//   GET  /task/                - Provides documentation
//   GET  /task/?name=[name]    - Gets a new task token
//   POST /task/[token]         - Post form data in reply
//   GET  /task/square/         - Gets documentation
//   GET  /task/square/[token]  - Gets a number to square
//   POST /task/square/[token]  - Post your squared number
//
//   GET  /status         - HTML showing status
//   GET  /status/updates - Websocket providing updates to status
//
package main

import "io"

func writeDocumentation(wr io.Writer) {
	wr.Write([]byte("Documentation!"))
}
