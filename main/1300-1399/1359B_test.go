package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1359B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 10 1
.
1 2 10 1
..
2 1 10 1
.
.
3 3 3 7
..*
*..
.*.
outputCopy
10
1
20
18`
	testutil.AssertEqualCase(t, rawText, 0, CF1359B)
}
