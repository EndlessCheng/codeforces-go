package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1194D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0 3
3 3
3 4
4 4
outputCopy
Bob
Alice
Bob
Alice`
	testutil.AssertEqualCase(t, rawText, 0, CF1194D)
}
