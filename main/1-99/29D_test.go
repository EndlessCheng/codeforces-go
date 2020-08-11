package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF29D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2
2 3
3
outputCopy
1 2 3 2 1 
inputCopy
6
1 2
1 3
2 4
4 5
4 6
5 6 3
outputCopy
1 2 4 5 4 6 4 2 1 3 1 
inputCopy
6
1 2
1 3
2 4
4 5
4 6
5 3 6
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF29D)
}
