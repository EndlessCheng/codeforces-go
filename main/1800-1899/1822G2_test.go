package main

import (
	"bufio"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/contest/1822/problem/G2
// https://codeforces.com/problemset/status/1822/problem/G2
func TestCF1822G2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
5
1 7 7 2 7
3
6 2 18
9
1 2 3 4 5 6 7 8 9
4
1000 993 986 179
7
1 10 100 1000 10000 100000 1000000
8
1 1 2 2 4 4 8 8
9
1 1 1 2 2 2 4 4 4
outputCopy
6
1
3
0
9
16
45
inputCopy
7
20
72 52 22 91 64 37 8 54 14 17 80 56 21 74 89 16 72 93 22 30
20
49 40 4 80 3 97 99 72 21 53 89 24 80 58 2 61 21 66 19 33
20
53 53 51 98 30 95 53 64 36 44 71 42 18 98 92 14 61 74 88 56
20
29 99 23 11 68 18 53 49 61 28 45 78 42 79 49 84 62 2 4 90
20
48 80 48 12 22 45 52 65 7 25 40 10 81 43 67 56 99 39 76 37
20
32 44 42 100 51 31 83 16 89 62 2 37 47 70 81 68 53 1 87 61
20
95 9 63 56 76 55 47 78 98 53 54 79 95 97 71 3 91 83 67 6
outputCopy
0
0
6
0
0
0
0
inputCopy
1
20
32 44 42 100 51 31 83 16 89 62 2 37 47 70 81 68 53 1 87 61
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, CF1822G2)
}

func CF1822G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int64{}
		for ; n > 0; n-- {
			Fscan(in, &v)
			cnt[v]++
		}
		ans := int64(0)
		for v, c := range cnt {
			ans += c * (c - 1) * (c - 2)
			for p := 2; p*p <= v; p++ {
				if v%(p*p) == 0 {
					ans += c * cnt[v/p] * cnt[v/(p*p)]
				}
			}
		}
		Fprintln(out, ans)
	}
}

func TestCompareCF1822G2(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(3,1000)
		rg.NewLine()
		rg.IntSlice(n/2, 1,100)
		rg.IntSlice(n-n/2, 1e8,1e9)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, CF1822G, CF1822G2)
}
