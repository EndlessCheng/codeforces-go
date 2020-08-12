package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1003F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
to be or not to be
outputCopy
12
inputCopy
10
a ab a a b ab a a b c
outputCopy
13
inputCopy
6
aa bb aa aa bb bb
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1003F)
}
