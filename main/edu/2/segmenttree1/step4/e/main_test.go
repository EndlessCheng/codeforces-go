package main

import (
	"bytes"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 9
1 0 3
1 2 5
2 0 4 3
1 1 4
1 2 7
2 1 3 6
1 3 8
1 4 4
2 0 5 10
outputCopy
1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

func Test2(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		buf := &bytes.Buffer{}
		var n, q int
		n = 1e5
		q = 1e5
		buf.WriteString(strconv.Itoa(n) + " " + strconv.Itoa(q) + "\n")
		for i := 0; i < q; i++ {
			op := rand.Intn(2) + 1
			buf.WriteString(strconv.Itoa(op) + " ")
			if op == 1 {
				i, v := rand.Intn(n), rand.Intn(10)+1
				buf.WriteString(strconv.Itoa(i) + " " + strconv.Itoa(v))
			} else {
				l := rand.Intn(n)
				r := rand.Intn(n-l) + l + 1
				upp := rand.Intn(10) + 1
				buf.WriteString(strconv.Itoa(l) + " " + strconv.Itoa(r) + " " + strconv.Itoa(upp))
			}
			buf.WriteByte('\n')
		}
		buf.WriteByte('\n')
		//Println(buf.String())
		return buf.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		return
		var n, q, op, i, v, l, r, upp int
		Fscan(in, &n, &q)
		a := make([]int, n)
		for ; q > 0; q-- {
			Fscan(in, &op)
			if op == 1 {
				Fscan(in, &i, &v)
				a[i] = v
			} else {
				Fscan(in, &l, &r, &upp)
				cnt := 0
				for i := l; i < r; i++ {
					if 0 < a[i] && a[i] <= upp {
						cnt++
						a[i] = 0
					}
				}
				Fprintln(out, cnt)
			}
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
