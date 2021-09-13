package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1108/D
// https://codeforces.com/problemset/status/1108/problem/D
func TestCF1108D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
RBGRRBRGG
outputCopy
2
RBGRGBRGR
inputCopy
8
BBBGBRRR
outputCopy
2
BRBGBRGR
inputCopy
13
BBRRRRGGGGGRR
outputCopy
6
BGRBRBGBGBGRG`
	testutil.AssertEqualCase(t, rawText, 0, CF1108D)
}
