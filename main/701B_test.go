package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF701B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 1
3 1
2 2
outputCopy
4 2 0 
inputCopy
5 2
1 5
5 1
outputCopy
16 9 
inputCopy
100000 1
300 400
outputCopy
9999800001 `
	testutil.AssertEqualCase(t, rawText, 0, CF701B)
}
