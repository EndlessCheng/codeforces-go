package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1483/B
// https://codeforces.com/problemset/status/1483/problem/B
func TestCF1483B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
5 9 2 10 15
6
1 2 4 2 4 2
2
1 2
1
1
1
2
outputCopy
2 2 3 
2 2 1 
2 2 1 
1 1 
0 
inputCopy
16
7
13 13 9 15 7 8 14
24
18 18 2 5 20 7 9 6 9 17 14 7 5 10 17 3 13 3 13 16 15 3 8 16
40
8 5 1 13 16 12 1 20 14 8 6 1 7 15 15 2 1 6 14 9 3 7 11 4 10 16 13 16 16 2 3 18 18 4 13 3 18 6 10 6
9
10 14 10 4 19 6 9 15 14
5
20 11 4 5 8
11
19 12 7 4 7 6 13 14 6 3 4
3
6 8 19
5
20 3 7 16 4
16
13 15 5 17 1 9 9 6 4 5 6 12 9 13 13 18
10
3 11 12 9 3 19 8 10 4 10
18
20 14 13 11 18 15 8 15 13 4 9 17 17 7 18 20 4 15
4
10 17 9 1
1
3
3
17 19 16
7
17 11 3 12 2 12 15
4
10 17 12 1
outputCopy
5 3 5 1 4 2 
15 4 6 10 13 15 17 19 21 23 7 11 14 20 24 12 
25 2 4 7 12 14 16 18 20 22 24 27 31 35 3 13 17 21 25 36 19 26 23 28 29 30 
2 5 9 
2 2 4 
10 2 4 6 8 11 3 7 1 9 5 
1 3 
3 2 4 3 
11 2 4 6 10 14 16 3 7 15 5 1 
7 2 6 1 7 8 9 10 
16 3 5 7 9 11 14 18 4 10 15 6 12 16 8 17 13 
3 2 4 3 
0 
2 2 1 
2 2 1 
2 2 4 
inputCopy
1
4
10 17 9 1
outputCopy
3 2 4 3
inputCopy
1
6
8 19 10 14 7 10
outputCopy
4 2 6 1 3`
	testutil.AssertEqualCase(t, rawText, -1, CF1483B)
}

func TestCompareCF1483B(t *testing.T) {
	return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 20)
		rg.NewLine()
		rg.IntSlice(n, 1, 300)
		return rg.String()
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var T, n int
		for Fscan(in, &T); T > 0; T-- {
			Fscan(in, &n)
			type pair struct{ v,i int }
			a := make([]pair, n)
			for i := range a {
				Fscan(in, &a[i].v)
				a[i].i = i+1
			}
			ans := []int{}
			for len(a) > 0{
				b := a
				b = append(b, b[0])
				a = nil
				keep := false
				for i := 1; i < len(b); i++ {
					a = append(a, b[i-1])
					if gcd(b[i-1].v, b[i].v) == 1 {
						ans = append(ans, b[i].i)
						keep = true
						if i == len(b)-1 {
							a = a[1:]
							break
						}
						i++
					}
				}
				//Println(a)
				if !keep {
					break
				}
			}
			Fprint(out, len(ans)," ")
			for _, v := range ans {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1483B)
}
