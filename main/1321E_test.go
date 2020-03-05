package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1321E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3 3
2 3
4 7
2 4
3 2
5 11
1 2 4
2 1 6
3 4 6
outputCopy
1
inputCopy
10 10 10
51 87
68 743
82 1284
73 936
55 242
89 1554
58 357
66 666
62 512
51 87
92 1670
56 280
56 280
91 1631
77 1091
89 1554
61 473
98 1901
77 1091
77 1091
86 15 386
43 58 386
68 33 386
39 62 386
95 6 386
74 27 386
94 7 386
17 84 386
97 4 386
45 56 386
outputCopy
212`
	testutil.AssertEqualCase(t, rawText, 0, CF1321E)
}
