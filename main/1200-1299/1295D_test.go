package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1295D(t *testing.T) {
	// just copy from website
	rawText := `
3
4 9
5 10
42 9999999967
outputCopy
6
1
9999999966`
	testutil.AssertEqualCase(t, rawText, 0, CF1295D)
}
