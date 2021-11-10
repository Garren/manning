package main

import "testing"

func TestNameIt(t *testing.T) {
	name := getName()
	if name != "World!" {
		t.Error("Response from getName is unexpected value")
	}
}
