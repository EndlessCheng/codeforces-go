package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSolP3864(t *testing.T) {
	inputs := []string{
		`4734
NMSL
GREG
LSDC`,
	}
	answers := []string{
		`GREG`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 0, SolP3864)
}
