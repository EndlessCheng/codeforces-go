package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1368E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 6
1 2
1 3
2 3
2 4
3 4
3 4
7 6
1 2
1 3
2 4
2 5
3 6
3 7
outputCopy
2
3 4 
4
4 5 6 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF1368E)
}
