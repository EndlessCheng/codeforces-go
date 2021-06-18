package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/555/C
// https://codeforces.com/problemset/status/555/problem/C
func TestCF555C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 5
3 4 U
6 1 L
2 5 L
1 6 U
4 3 U
outputCopy
4
3
2
1
2
inputCopy
10 6
2 9 U
10 1 U
1 10 U
8 3 L
10 1 L
6 5 U
outputCopy
9
1
10
6
0
2
inputCopy
204 10
179 26 U
176 29 L
99 106 U
44 161 L
88 117 U
165 40 U
116 89 U
118 87 U
62 143 U
26 179 U
outputCopy
26
176
77
44
88
11
60
58
114
18`
	testutil.AssertEqualCase(t, rawText, -1, CF555C)
}
