package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSolP1204(t *testing.T) {
	inputs := []string{
		`3
300 1000
700 1200
1500 2100`, `10
2 3
4 5
6 7
8 9
10 11
12 13
14 15
16 17
18 19
1 20`,
	}
	answers := []string{
		`900 300`, `19 0`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 0, SolP1204)
}
