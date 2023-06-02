package main

import (
	"bufio"
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math"
	"testing"
)

// https://codeforces.com/problemset/problem/1272/F
// https://codeforces.com/problemset/status/1272/problem/F
func TestCF1272F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
(())(()
()))()
outputCopy
(())()()
inputCopy
)
((
outputCopy
(())
inputCopy
)
)))
outputCopy
((()))
inputCopy
())
(()(()(()(
outputCopy
(()()()(()()))
inputCopy
)(())()(()()(())())()))()))))))()))((()((()))(()(()(((()()(()))()()()()()((()(()()(())()))))())())())()(())()())(())(()()(()())()()())))(())((()))))()())((()))))()()(()()))))))()()(()()()(()()()()())
)())))))())())()((((((()())())()()((()(())((())(())((()(((()))(((()()())))))(((()()(())()()((())()()))(()())(())(((()())))))(())(((()()))((((()))(()()))((())()(()(()()(()(((()))()))()()()())())()(())
outputCopy
(()(())()(()()()())())()(((((()()())())()()((()()())((()()(())((()(()(()))(((()()()))()))()()()()()(()()()()()(())()()))(()())())())()(())()())(())(()()(())())()()()()))(())((()))()()()())((())()(()(()()(()(()()))()))()()()()()()()()()()()()())
inputCopy
(())(
)()((
outputCopy
()()()()`
	testutil.AssertEqualCase(t, rawText, 0, CF1272F)
}

func TestCompare(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		return `(())(
)()((`
		rg := testutil.NewRandGenerator()
		rg.StrInSet(1,5, "()")
		rg.NewLine()
		rg.StrInSet(1,5, "()")
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, solveCF1272F, CF1272F)
}

func solveCF1272F(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var s, t string
	fmt.Fscan(bufio.NewReader(in), &s, &t)
	n := len(s)
	m := len(t)
	dp := makeGrid(n+1, m+1)
	const (
		BracketMax = 201
		Inf        = math.MaxInt32
	)
	for spos := 0; spos < n+1; spos++ {
		for tpos := 0; tpos < m+1; tpos++ {
			for i := 0; i < BracketMax; i++ {
				dp[spos][tpos][i][0] = Inf
				dp[spos][tpos][i][1] = Inf
			}
		}
	}
	dp[0][0][0][0] = 0
	for spos := 0; spos < n+1; spos++ {
		for tpos := 0; tpos < m+1; tpos++ {
			for i := 0; i < BracketMax-1; i++ {
				for j := 0; j < 2; j++ {
					if dp[spos][tpos][i][j] == Inf {
						continue
					}
					switch {
					case spos == n && tpos < m:
						if t[tpos] == '(' {
							dp[spos][tpos+1][i+1][0] = min(dp[spos][tpos+1][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos][tpos][i-1][1] = min(dp[spos][tpos][i-1][1], dp[spos][tpos][i][j]+1)
							}
						} else {
							dp[spos][tpos][i+1][0] = min(dp[spos][tpos][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos][tpos+1][i-1][1] = min(dp[spos][tpos+1][i-1][1], dp[spos][tpos][i][j]+1)
							}
						}
					case spos < n && tpos == m:
						if s[spos] == '(' {
							dp[spos+1][tpos][i+1][0] = min(dp[spos+1][tpos][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos][tpos][i-1][1] = min(dp[spos][tpos][i-1][1], dp[spos][tpos][i][j]+1)
							}
						} else {
							dp[spos][tpos][i+1][0] = min(dp[spos][tpos][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos+1][tpos][i-1][1] = min(dp[spos+1][tpos][i-1][1], dp[spos][tpos][i][j]+1)
							}
						}
					case spos < n && tpos < m:
						switch {
						case s[spos] == '(' && t[tpos] == '(':
							dp[spos+1][tpos+1][i+1][0] = min(dp[spos+1][tpos+1][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos][tpos][i-1][1] = min(dp[spos][tpos][i-1][1], dp[spos][tpos][i][j]+1)
							}
						case s[spos] == '(' && t[tpos] == ')':
							dp[spos+1][tpos][i+1][0] = min(dp[spos+1][tpos][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos][tpos+1][i-1][1] = min(dp[spos][tpos+1][i-1][1], dp[spos][tpos][i][j]+1)
							}
						case s[spos] == ')' && t[tpos] == '(':
							dp[spos][tpos+1][i+1][0] = min(dp[spos][tpos+1][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos+1][tpos][i-1][1] = min(dp[spos+1][tpos][i-1][1], dp[spos][tpos][i][j]+1)
							}
						case s[spos] == ')' && t[tpos] == ')':
							dp[spos][tpos][i+1][0] = min(dp[spos][tpos][i+1][0], dp[spos][tpos][i][j]+1)
							if i-1 >= 0 {
								dp[spos+1][tpos+1][i-1][1] = min(dp[spos+1][tpos+1][i-1][1], dp[spos][tpos][i][j]+1)
							}
						}
					}
				}
			}
		}
	}
	for i := BracketMax - 1; i > 0; i-- {
		for j := 0; j < 2; j++ {
			if dp[n][m][i][j] != Inf {
				dp[n][m][i-1][1] = min(dp[n][m][i-1][1], dp[n][m][i][j]+1)
			}
		}
	}
	length := dp[n][m][0][1]
	ans := make([]int, length)
	var (
		dfs func(spos, tpos, bracket, pos int) bool
	)
	dfs = func(spos, tpos, bracket, pos int) bool {
		if pos == 0 {
			return true
		}
		for i := 0; i < 2; i++ {
			if bracket+i*2-1 < 0 {
				continue
			}
			if dp[spos][tpos][bracket][i] == pos {
				pos--
				ans[pos] = i
				if dfs(spos, tpos, bracket+i*2-1, pos) {
					return true
				}
				switch {
				case spos > 0 && tpos == 0:
					if s[spos-1] == "()"[i] && dfs(spos-1, tpos, bracket+i*2-1, pos) {
						return true
					}
				case spos == 0 && tpos > 0:
					if t[tpos-1] == "()"[i] && dfs(spos, tpos-1, bracket+i*2-1, pos) {
						return true
					}
				case spos > 0 && tpos > 0:
					if s[spos-1] == "()"[i] && dfs(spos-1, tpos, bracket+i*2-1, pos) {
						return true
					}
					if t[tpos-1] == "()"[i] && dfs(spos, tpos-1, bracket+i*2-1, pos) {
						return true
					}
					if t[tpos-1] == "()"[i] && s[spos-1] == "()"[i] && dfs(spos-1, tpos-1, bracket+i*2-1, pos) {
						return true
					}
				}
			}
		}
		return false
	}
	dfs(n, m, 0, length)
	//fmt.Fprintln(out,length)
	for i := 0; i < length; i++ {
		fmt.Fprintf(out, "%c", "()"[ans[i]])
	}
}
func makeGrid(h, w int) [][][201][2]int {
	index := make([][][201][2]int, h, h)
	data := make([][201][2]int, h*w, h*w)
	for i := 0; i < h; i++ {
		index[i] = data[i*w : (i+1)*w]
	}
	return index
}