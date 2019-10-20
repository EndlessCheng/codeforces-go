package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSolP1202(t *testing.T) {
	inputs := []string{
		`20`,
	}
	answers := []string{
		`36 33 34 33 35 35 34`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 0, SolP1202)
}
