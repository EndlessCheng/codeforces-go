package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1341A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
7 20 3 101 18
11 11 10 234 2
8 9 7 250 122
19 41 21 321 10
3 10 8 6 1
outputCopy
Yes
No
Yes
No
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1341A)
}
