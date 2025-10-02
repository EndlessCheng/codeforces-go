package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf348C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, k, x int
	var op string
	Fscan(in, &n, &m, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	B := int(math.Sqrt(1e5 * 3))
	sets := make([][]int, m)
	iToBig := make([]int, m)
	bigSum := []int{}
	for i := range sets {
		Fscan(in, &k)
		sets[i] = make([]int, k)
		s := 0
		for j := range sets[i] {
			Fscan(in, &sets[i][j])
			sets[i][j]--
			s += a[sets[i][j]]
		}
		if k > B {
			bigSum = append(bigSum, s)
			iToBig[i] = len(bigSum)
		}
	}

	nb := len(bigSum)
	intersectionSize := make([][]int32, m)
	for i := range intersectionSize {
		intersectionSize[i] = make([]int32, nb)
	}
	has := make([]bool, n)
	cur := 0
	for _, big := range sets {
		if len(big) <= B {
			continue
		}
		for _, j := range big {
			has[j] = true
		}
		for i, st := range sets {
			for _, j := range st {
				if has[j] {
					intersectionSize[i][cur]++
				}
			}
		}
		cur++
		for _, j := range big {
			has[j] = false
		}
	}

	bigTodo := make([]int, nb)
	for range q {
		Fscan(in, &op, &k)
		k--
		st := sets[k]
		if op == "+" {
			Fscan(in, &x)
			if len(st) > B {
				bigTodo[iToBig[k]-1] += x
			} else {
				for _, i := range st {
					a[i] += x
				}
				for j, sz := range intersectionSize[k] {
					bigSum[j] += x * int(sz)
				}
			}
		} else {
			s := 0
			if len(st) > B {
				s = bigSum[iToBig[k]-1]
			} else {
				for _, i := range st {
					s += a[i]
				}
			}
			for j, sz := range intersectionSize[k] {
				s += bigTodo[j] * int(sz)
			}
			Fprintln(out, s)
		}
	}
}

//func main() { cf348C(bufio.NewReader(os.Stdin), os.Stdout) }
