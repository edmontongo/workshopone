package main

import (
	"fmt"
	"time"
)

type SysIntError uint8

func (s SysIntError) IsFatal() bool { return s > 3 }

func (s SysIntError) Error() string {
	switch s {
	case 1:
		return "system could not find file"
	case 2:
		return "floating point overflow"
	case 3:
		return "graphics card overheating"
	case 4:
		return "PSU issuing smoke"
	case 5:
		return "CPU meltdown"
	}
	return ""
}

var globalVariable int // Psst, hide this from students. BAD practice !!!
func RenderGraphics() error {
	globalVariable++
	if globalVariable > 3 {
		return nil
	}
	return SysIntError(globalVariable)
}
func main() {
	for attempt := 1; attempt <= 5; attempt++ {
		err := RenderGraphics() // HL
		if err == nil {         // HL
			fmt.Println("SUCCESS: Rendering finished")
			break // HL
		}
		fmt.Printf("FAIL: attempt #%d: %s\n", attempt, err)
		time.Sleep(250 * time.Millisecond)
	}
}
