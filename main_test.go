package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(5, 6) != 11 {
		t.Fail()
	}

	if Add(5, 5) == 11 {
		t.Fail()
	}
}
