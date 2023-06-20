package main

import (
	"math"
	"testing"
)

func TestDefault(t *testing.T) {
	got := math.Abs(-1)

	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func TestChangeName(t *testing.T) {
	_, err := changeName("Daniel", "Hendrik")

	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestSameName(t *testing.T) {
	_, err := changeName("Daniel", "daniel")

	if err == nil {
		t.Errorf("same name not detected")
		return
	}
}
func TestMissingName(t *testing.T) {
	_, err := changeName("Daniel", "")

	if err == nil {
		t.Errorf("missing name not detected")
		return
	}
}
