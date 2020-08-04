package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

func TestCF721D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3 1
5 4 3 5 2
outputCopy
5 4 3 5 -1 
inputCopy
5 3 1
5 4 3 5 5
outputCopy
5 4 0 5 5 
inputCopy
5 3 1
5 4 4 5 5
outputCopy
5 1 4 5 5 
inputCopy
3 2 7
5 4 2
outputCopy
5 11 -5 
inputCopy
3 2 7
5 4 2
outputCopy
5 11 -5 `
	testutil.AssertEqualCase(t, rawText, 0, CF721D)
}

// 无尽对拍
func TestCF721DCmp(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
//		return `2 3 10
//-2 -5`
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		rg.Int(1, 3)
		rg.Int(1, 10)
		rg.NewLine()
		rg.IntSlice(n, -10, 10)
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var n, k, x int
		Fscan(in, &n, &k, &x)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := int(1e18)
		b := make([]int, n)
		var f func(p, left int)
		f = func(p, left int) {
			//println(p,left)
			if left == 0 {
				s := 1
				for _, v := range a {
					s *= v
				}
				if s < ans {
					ans = s
					copy(b, a)
				}
			}
			if p == n {
				return
			}
			for i := 0; i <= left; i++ {
				a[p]+=i*x
				f(p+1,left-i)
				a[p]-=2*i*x
				f(p+1,left-i)
				a[p]+=i*x
			}
		}
		f(0, k)
		Fprint(out, ans)
		//for _, v := range b {
		//	Fprint(out, v, " ")
		//}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF721D)
}
