package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
const mod = 1_000_000_007

func run(in io.Reader, out io.Writer) {
	var n, k, op int
	Fscan(in, &n, &k, &op)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	if k == 1 {
		slices.Sort(a)
		for _, v := range a {
			Fprintln(out, v)
		}
		return
	}

	mx := slices.Max(a)
	h := hp{a}
	heap.Init(&h)
	for ; op > 0 && a[0] < mx; op-- {
		a[0] *= k
		heap.Fix(&h, 0)
	}

	slices.Sort(a)
	pm := pow(k, op/n)
	for _, v := range a[op%n:] {
		Fprintln(out, v%mod*pm%mod)
	}
	pm = pm * k % mod
	for _, v := range a[:op%n] {
		Fprintln(out, v%mod*pm%mod)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

type hp struct{ sort.IntSlice }
func (hp) Push(any)     {}
func (hp) Pop() (_ any) { return }
func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
