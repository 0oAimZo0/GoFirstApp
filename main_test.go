package main

import "testing"

func Test(t *testing.T) {
	a := Plus(1, 2)
	if a != 3 {
		t.Errorf("eiei")
	}
}
