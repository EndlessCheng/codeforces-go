package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1140C(t *testing.T) {
	// just copy from website
	rawText := `
4 3
4 7
15 1
3 6
6 8
outputCopy
78
inputCopy
5 3
12 31
112 4
100 100
13 55
55 50
outputCopy
10000`
	testutil.AssertEqual(t, rawText, CF1140C)
}
