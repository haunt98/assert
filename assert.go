package assert

import "log"

// Bool check got and want
func Bool(want bool, got bool, msg string) {
	if got != want {
		log.Fatalf("assert error: %s", msg)
	}
}

// True check got is true
func True(got bool, msg string) {
	Bool(true, got, msg)
}

// False check got is false
func False(got bool, msg string) {
	Bool(false, got, msg)
}
