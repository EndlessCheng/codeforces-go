package main

import (
	"bufio"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	s := "80\n"
	for x := -4; x <= 4; x++ {
		for y := -4; y <= 4; y++ {
			if x != 0 || y != 0 {
				s += Sprintln(x, y)
			}
		}
	}
	inputs := []string{s}
	//testutil.AssertEqualStringCase(t, testCases, 0, run)
	testutil.AssertEqualRunResults(t, inputs, 0, runAC, run)
}

func runAC(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type tuple struct {
		x, y int
		s    string
	}

	ans := map[[2]int]string{}
	for q := []tuple{{}}; len(q[0].s) <= 3; {
		p := q[0]
		q = q[1:]
		if ans[[2]int{p.x, p.y}] == "" {
			ans[[2]int{p.x, p.y}] = p.s
		}
		q = append(q,
			tuple{p.x + 1<<len(p.s), p.y, p.s + "E"},
			tuple{p.x - 1<<len(p.s), p.y, p.s + "W"},
			tuple{p.x, p.y + 1<<len(p.s), p.s + "N"},
			tuple{p.x, p.y - 1<<len(p.s), p.s + "S"},
		)
	}

	var t int
	var p [2]int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: ", _case)
		Fscan(in, &p[0], &p[1])
		if s, ok := ans[p]; ok {
			Fprintln(out, s)
		} else {
			Fprintln(out, "IMPOSSIBLE")
		}
	}
}
