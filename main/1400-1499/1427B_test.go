package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1427B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
5 2
WLWLL
6 5
LLLWWL
7 1
LWLWLWL
15 5
WWWLLLWWWLLLWWW
40 7
LLWLWLWWWLWLLWLWWWLWLLWLLWLLLLWLLWWWLWWL
1 0
L
1 1
L
6 1
WLLWLW
outputCopy
7
11
6
26
46
0
1
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1427B)
}
