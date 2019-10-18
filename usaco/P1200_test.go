package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSolP1200(t *testing.T) {
	inputs := []string{
		`COMETQ
HVNGAT`,
		`ABSTAR
USACO`,
	}
	answers := []string{
		`GO`, `STAY`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 0, SolP1200)
}
