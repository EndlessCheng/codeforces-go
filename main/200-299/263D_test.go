package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF263D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 2
1 2
2 3
3 1
outputCopy
3
1 2 3 
inputCopy
4 6 3
4 3
1 2
1 3
1 4
2 3
2 4
outputCopy
4
3 4 1 2 
inputCopy
7 8 2
1 2
2 3
3 4
4 2
1 5
5 6
6 7
7 5
outputCopy
3
5 6 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF263D)
}
