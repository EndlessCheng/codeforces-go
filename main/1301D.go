package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1301D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k int
	Fscan(in, &n, &m, &k)
	if k > 4*n*m-2*n-2*m {
		Fprint(out, "NO")
		return
	}

	type pair struct {
		times int
		op    string
	}
	ans := []pair{}
	makeAns := func(ops []string, maxTimes, left int) int {
		for _, op := range ops {
			if left <= len(op)*maxTimes {
				if left > 0 {
					if len(op) == 1 {
						ans = append(ans, pair{left, op})
					} else {
						a, b := left/len(op), left%len(op)
						if a > 0 {
							ans = append(ans, pair{a, op})
						}
						if b > 0 {
							ans = append(ans, pair{1, op[:b]})
						}
					}
				}
				return 0
			}
			ans = append(ans, pair{maxTimes, op})
			left -= len(op) * maxTimes
		}
		return left
	}
	printAns := func() {
		Fprintln(out, "YES")
		Fprintln(out, len(ans))
		for _, a := range ans {
			Fprintln(out, a.times, a.op)
		}
	}

	if n == 1 {
		makeAns([]string{"R", "L"}, m-1, k)
		printAns()
		return
	}
	var cycleNum, left int
	if cycleSize := 4*(n-1) + 1; k <= cycleSize*(m-1) {
		cycleNum, left = k/cycleSize, k%cycleSize
		k = 0
	} else {
		cycleNum, left = m-1, 0
		k -= cycleSize * (m - 1)
	}
	for i := 0; i < cycleNum; i++ {
		ans = append(ans, pair{n - 1, "D"}, pair{n - 1, "RLU"}, pair{1, "R"})
	}
	makeAns([]string{"D", "RLU"}, n-1, left)
	k = makeAns([]string{"D", "U"}, n-1, k)
	if k > 0 {
		ans = append(ans, pair{k, "L"})
	}
	printAns()
}

//func main() { CF1301D(os.Stdin, os.Stdout) }
