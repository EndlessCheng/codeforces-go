package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1493/problem/B
// https://codeforces.com/problemset/status/1493/problem/B
func TestCF1493B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
24 60
12:21
24 60
23:59
90 80
52:26
1 100
00:01
10 10
04:04
outputCopy
12:21
00:00
52:28
00:00
00:00
inputCopy
15
1 1
00:00
1 100
00:01
100 100
02:82
10 10
04:04
100 100
99:00
100 2
30:00
50 3
30:00
51 3
30:01
52 16
21:12
100 20
20:07
23 100
21:45
11 11
02:02
89 89
88:86
100 1
01:00
1 100
00:01
outputCopy
00:00
00:00
02:82
00:00
00:00
00:00
00:00
50:00
21:12
20:08
21:50
10:00
88:88
00:00
00:00`
	testutil.AssertEqualCase(t, rawText, 0, CF1493B)
}
