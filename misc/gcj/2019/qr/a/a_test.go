package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	inputs := []string{
		`
3
4
940
4444`,
	}
	answers := []string{
		`
Case #1: 2 2
Case #2: 852 88
Case #3: 667 3777`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 0, run)
}
