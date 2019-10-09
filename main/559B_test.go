package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol559B(t *testing.T) {
	// just copy from website
	rawText := `
aaba
abaa
outputCopy
YES
inputCopy
aabb
abab
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, -1, Sol559B)
}
