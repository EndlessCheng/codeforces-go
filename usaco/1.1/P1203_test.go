package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSolP1203(t *testing.T) {
	inputs := []string{
		`29 
wwwbbrwrbrbrrbrbrwrwwrbwrwrrb`,
		`77
rwrwrwrwrwrwrwrwrwrwrwrwbwrwbwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwrwr`,
	}
	answers := []string{
		`11`, `74`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 2, SolP1203)
}
