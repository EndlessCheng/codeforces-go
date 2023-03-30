package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/contest/1374/problem/E2
// https://codeforces.com/problemset/status/1374/problem/E2
func TestCF1374E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3 1
6 0 0
11 1 0
9 0 1
21 1 1
10 1 0
8 0 1
outputCopy
24
6 5 1 
inputCopy
6 3 2
6 0 0
11 1 0
9 0 1
21 1 1
10 1 0
8 0 1
outputCopy
39
4 6 5 
inputCopy
4 1 1
1 0 0
4 0 1
1 0 1
5 1 0
outputCopy
-1
inputCopy
5 2 2
4 1 1
5 1 1
3 0 1
2 1 0
1 1 0
outputCopy
9
1 2
inputCopy
3 2 1
2 1 1
4 0 1
3 1 1
outputCopy
5
1 3
inputCopy
4 3 1
4 0 0
1 1 0
5 0 1
1 1 1
outputCopy
6
inputCopy
6 4 1
4 1 0
3 1 1
2 0 1
1 1 0
1 0 1
2 1 1
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1374E2)
}

func TestCompare(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 9)
		m := rg.Int(1, n)
		k := rg.Int(1, m)
		_ = k
		rg.NewLine()
		for i := 0; i < n; i++ {
			rg.Int(1, 5)
			rg.Int(0, 1)
			rg.Int(0, 1)
			rg.NewLine()
		}
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runAC_CF1374E2, CF1374E2)
}


func TestCheckCF1374E2(_t *testing.T) {
	//return
	assert := assert.New(_t)
	_ = assert

	testutil.DebugTLE = 0

	inputGenerator := func() (string, testutil.OutputChecker) {
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 5)
		m := rg.Int(1, n)
		k := rg.Int(1, m)
		_ = k
		rg.NewLine()
		for i := 0; i < n; i++ {
			rg.Int(1, 5)
			rg.Int(0, 1)
			rg.Int(0, 1)
			rg.NewLine()
		}
		return rg.String(), func(myOutput string) (_b bool) {
			// 检查 myOutput 是否符合题目要求
			// * 最好重新看一遍题目描述以免漏判 *
			// 对于 special judge 的题目，可能还需要额外跑个暴力来检查 myOutput 是否满足最优解等
			//in := strings.NewReader(myOutput)
			//
			//myA := make([]int, n)
			//for i := range myA {
			//	Fscan(in, &myA[i])
			//}
			//if !assert.EqualValues(a, myA) {
			//	return
			//}

			return true
		}
	}

	target := 0
	testutil.CheckRunResultsInfWithTarget(_t, inputGenerator, target, CF1374E2)
}

const MAX = int64(1 << 60)

type data struct {
	t, i int
	s    int64
}

type Data []data

func (p Data) Len() int           { return len(p) }
func (p Data) Less(i, j int) bool { return p[i].t < p[j].t }
func (p Data) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Data) IndexValueS(i int) int64 {
	if i < 0 || i >= len(p) {
		return 0
	}

	return p[i].s
}

func (p Data) IndexValueT(i int) int {
	if i < 0 || i >= len(p) {
		return 1 << 30
	}

	return p[i].t
}

func runAC_CF1374E2(in io.Reader, out io.Writer) {
	var n, m, k int
	datas := make([]Data, 4)

	fmt.Fscanf(in, "%d %d %d\n", &n, &m, &k)

	for i := 0; i < 4; i++ {
		datas[i] = make(Data, 0, n)
	}

	var t, a, b int
	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%d %d %d\n", &t, &a, &b)

		datas[b<<1+a] = append(datas[b<<1+a], data{t: t, i: i, s: int64(t)})
	}

	// if len(aa)+len(cc) < k || len(bb)+len(cc) < k ||
	// 	2*k-len(cc) > m {
	// 	fmt.Println(-1)
	// 	return
	// }

	sort.Sort(datas[0])
	sort.Sort(datas[1])
	sort.Sort(datas[2])
	sort.Sort(datas[3])

	for _, d := range datas {
		for j := 1; j < len(d); j++ {
			d[j].s += d[j-1].s
		}
	}

	sum := MAX
	indexs := make([]int, 4)

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	f := func(cci int) int64 {
		returnSum := MAX

		nl := max(k-cci, 0)
		if nl > len(datas[1]) || nl > len(datas[2]) || cci+nl+nl > m {
			return returnSum
		}

		tempSum := datas[3].IndexValueS(cci-1) + datas[1].IndexValueS(nl-1) + datas[2].IndexValueS(nl-1)

		aai, bbi, ddi := nl, nl, 0

		for mm := m - (cci + nl + nl); mm > 0; mm-- {
			value, index := minIndex(datas[0].IndexValueT(ddi), datas[1].IndexValueT(aai), datas[2].IndexValueT(bbi))
			switch index {
			case 0:
				tempSum += value
				ddi++
			case 1:
				tempSum += value
				aai++
			case 2:
				tempSum += value
				bbi++
			}
		}

		if tempSum < sum {
			sum = tempSum
			indexs[0] = ddi
			indexs[1] = aai
			indexs[2] = bbi
			indexs[3] = cci
		}

		if tempSum < returnSum {
			returnSum = tempSum
		}

		return returnSum
	}

	left, right := 0, min(len(datas[3]), m)

	for left+3 <= right {
		lmid := left + (right-left)/3
		rmid := right - (right-left)/3

		if f(lmid) >= f(rmid) {
			left = lmid
		} else {
			right = rmid
		}
	}

	for ; left <= right; left++ {
		f(left)
	}

	if sum == MAX {
		fmt.Fprintln(out, -1)
		return
	}

	fmt.Fprintln(out, sum)
return
	ids := []int{}
	for i, d := range datas {
		maxIndex := indexs[i]
		for j := 0; j < maxIndex; j++ {
			ids = append(ids, d[j].i)
			//fmt.Fprint(out, d[j].i, " ")
		}
	}
	sort.Ints(ids)
	for _, v := range ids {
		fmt.Fprint(out, v, " ")
	}
}

func minIndex(n0, n1, n2 int) (int64, int) {
	if n0 <= n1 && n0 <= n2 {
		return int64(n0), 0
	}

	if n1 <= n0 && n1 <= n2 {
		return int64(n1), 1
	}

	if n2 <= n0 && n2 <= n1 {
		return int64(n2), 2
	}

	return 0, -1
}
