package main

import (
	"testing"
)

// Note: try using test tables
func TestIsArmstrong(t *testing.T) {
	// Test case 1
	if isArmstrong(153) != "Armstrong" {
		t.Errorf("Expected true, got false")
	}

	// Test case 2
	if isArmstrong(123) != "NotArmstrong" {
		t.Errorf("Expected false, got true")
	}

	// Test case 3
	if isArmstrong(9474) != "InputError" {
		t.Errorf("Expected InputError, got another result")
	}

	// Test case 4
	if isArmstrong(9475) != "InputError" {
		t.Errorf("Expected InputError, got another result")
	}
}
