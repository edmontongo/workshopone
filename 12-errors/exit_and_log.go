package main

import "log"

func init() {
	log.SetFlags(0)
}

// ThunderboltError is an imaginary error occurring on a ThunderBolt (TM) bus
type ThunderboltError int

func (t ThunderboltError) IsFatal() bool {
	return t < 0
}

func (t ThunderboltError) Error() string {
	switch t {
	case -1:
		return "handshake error"
	case 1:
		return "GPIO timing error"
	case 2:
		return "bus already in use"
	}
	return "unsupported error"
}

// An imaginary function that enables the thunderbolt 3 extension board
func KickThunderbolt() error {
	return ThunderboltError(-1)
}

// fatalist indicates whether an error is severe // HL
type fatalist interface {
	IsFatal() bool // HL
}

func main() {
	err := KickThunderbolt()
	if err != nil {
		if f, ok := err.(fatalist); ok && f.IsFatal() { // HL
			log.Fatalf("FATAL: could not connect display: %s", err)
		} else {
			log.Printf("PCI bus error: %s", err)
		}
	}
}
