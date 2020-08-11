package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF353B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
10 99
outputCopy
1
2 1 
inputCopy
2
13 24 13 45
outputCopy
4
1 2 2 1 
inputCopy
2
10 10 10 10
outputCopy
1
1 2 1 2
inputCopy
2
10 11 11 11
outputCopy
2
1 1 2 2`
	testutil.AssertEqualCase(t, rawText, 0, CF353B)
}
