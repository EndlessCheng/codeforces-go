package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1467/D
// https://codeforces.com/problemset/status/1467/problem/D
func TestCF1467D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 1 5
3 5 1 4 2
1 9
2 4
3 6
4 6
5 2
outputCopy
62
58
78
86
86
inputCopy
5 2 5
3 5 1 4 2
1 9
2 4
3 6
4 6
5 2
outputCopy
157
147
207
227
227
inputCopy
4 40 6
92 21 82 46
3 56
1 72
4 28
1 97
2 49
2 88
outputCopy
239185261
666314041
50729936
516818968
766409450
756910476`
	testutil.AssertEqualCase(t, rawText, 0, CF1467D)
}
