package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1334B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 3
5 1 2 1
4 10
11 9 11 9
2 5
4 3
3 7
9 4 9
outputCopy
2
4
0
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1334B)
}
