package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p4447(t *testing.T) {
	inputs := []string{
		`7
4 5 2 3 -4 -3 -5`,
	}
	answers := []string{
		`3`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 0, p4447)
}
