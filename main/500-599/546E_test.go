package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF546E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 2 6 3
3 5 3 1
1 2
2 3
3 4
4 2
outputCopy
YES
1 0 0 0 
2 0 0 0 
0 5 1 0 
0 0 2 1 
inputCopy
2 0
1 2
2 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF546E)
}
