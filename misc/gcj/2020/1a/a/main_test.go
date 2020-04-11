package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	customInputs := []string{
		`2
5
*CONUTS
*COCONUTS
*OCONUTS
*CONUTS
*S
2
*XZ
*XYZ`, `1
2
A*C*E
*B*D*`,
	}
	customAnswers := []string{
		`Case #1: COCONUTS
Case #2: *`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
}
