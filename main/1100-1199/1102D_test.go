package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1102/D
// https://codeforces.com/problemset/status/1102/problem/D
func TestCF1102D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
121
outputCopy
021
inputCopy
6
000000
outputCopy
001122
inputCopy
6
211200
outputCopy
211200
inputCopy
6
120110
outputCopy
120120`
	testutil.AssertEqualCase(t, rawText, 0, CF1102D)
}

func TestCompareCF1102D(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		k := rg.IntOnly(1,2)
		n := rg.Int(k*3,k*3)
		rg.NewLine()
		rg.Str(n, n,'0', '2')
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, solCF1102D, CF1102D)
}

func solCF1102D(in io.Reader, out io.Writer) {
	var n int
	var str []byte
	fmt.Fscan(in, &n, &str)
	cnt := make([]int,3)
	for i:= 0;i < n;i++{
		cnt[str[i]-byte('0')]++
	}
	cnt[0] -= n/3
	cnt[1] -= n/3
	cnt[2] -= n/3
	for i:= 0;i < n;i++{
		tmp := str[i]-byte('0')
		if cnt[tmp] > 0{
			for j := 0;j < int(tmp);j++{
				if cnt[j] < 0{
					str[i] = byte('0')+byte(j)
					cnt[j]++
					cnt[tmp]--
					break
				}
			}
		}
	}
	//	fmt.Printf("%s\n", str)
	for i:= n-1;i >= 0;i--{
		tmp := str[i]-byte('0')
		if cnt[tmp] > 0{
			for j := 2;j > int(tmp);j--{
				if cnt[j] < 0{
					str[i] = byte('0')+byte(j)
					cnt[j]++
					cnt[tmp]--
					break
				}
			}
		}
	}
	fmt.Fprintf(out, "%s", str)
}