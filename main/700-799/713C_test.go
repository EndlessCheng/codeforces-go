package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF713C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
2 1 5 11 5 9 11
outputCopy
9
inputCopy
5
5 4 3 2 1
outputCopy
12`
	testutil.AssertEqualCase(t, rawText, 0, CF713C)
}
