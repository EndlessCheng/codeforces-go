package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1334A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
3
0 0
1 1
1 2
2
1 0
1000 3
4
10 1
15 2
10 2
15 2
1
765 432
2
4 4
4 3
5
0 0
1 0
1 0
1 0
1 0
2
1 0
1 1
2
10 1
11 3
outputCopy
NO
YES
NO
YES
NO
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1334A)
}
