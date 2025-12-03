package main

import "testing"

func TestEventOrOdd(t *testing.T){
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("Expected even, actual %s", result)
	}
}