package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1208C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
8 9 1 13
3 12 7 5
0 2 4 11
6 10 15 14
inputCopy
8
outputCopy
19 55 11 39 32 36 4 52
51 7 35 31 12 48 28 20
43 23 59 15 0 8 16 44
3 47 27 63 24 40 60 56
34 38 6 54 17 53 9 37
14 50 30 22 49 5 33 29
2 10 18 46 41 21 57 13
26 42 62 58 1 45 25 61`
	testutil.AssertEqualCase(t, rawText, 0, CF1208C)
}
