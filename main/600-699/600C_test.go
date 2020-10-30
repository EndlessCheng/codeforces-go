package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF600C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aabc
outputCopy
abba
inputCopy
aabcd
outputCopy
abcba
inputCopy
aabbcccdd
outputCopy
abcdcdcba`
	testutil.AssertEqualCase(t, rawText, 0, CF600C)
}
