package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1328D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
1 2 1 2 2
6
1 2 2 1 2 2
5
1 2 1 2 3
3
10 10 10
3
1 2 1
outputCopy
2
1 2 1 2 2
2
2 1 2 1 2 1
3
2 3 2 3 1
1
1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1328D)
}
