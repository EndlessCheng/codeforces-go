package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1202B(t *testing.T) {
	// just copy from website
	rawText := `
0840
outputCopy
-1 17 7 7 7 -1 2 17 2 7 
17 17 7 5 5 5 2 7 2 7 
7 7 7 4 3 7 1 7 2 5 
7 5 4 7 3 3 2 5 2 3 
7 5 3 3 7 7 1 7 2 7 
-1 5 7 3 7 -1 2 9 2 7 
2 2 1 2 1 2 2 2 0 1 
17 7 7 5 7 9 2 17 2 3 
2 2 2 2 2 2 0 2 2 2 
7 7 5 3 7 7 1 3 2 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF1202B)
}
