package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1363C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
3 1
2 1
3 1
outputCopy
Ashish
inputCopy
1
3 2
1 2
1 3
outputCopy
Ayush`
	testutil.AssertEqualCase(t, rawText, 0, CF1363C)
}
