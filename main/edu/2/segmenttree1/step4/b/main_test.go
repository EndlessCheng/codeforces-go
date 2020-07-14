package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 4
0 1
0 0

2 1
1 2

0 0
0 2

1 0
0 2

1 4
2 3
1 3
2 2
outputCopy
0 2
0 0

0 2
0 1

0 1
0 0

2 1
1 2`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
