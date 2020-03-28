package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		// 排序+去重
		sort.Ints(a)
		j := 0
		for i := 1; i < n; i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		ans := a[j]
		n = j + 1
		a = a[:n]

		// 从大到小三重循环，内部的两层循环最多找 10 个非因子
		for i := n - 1; i >= 0; i-- {
			ai := a[i]
			cnt := 0
			for j := i - 1; j >= 0; j-- {
				aj := a[j]
				if ai%aj != 0 {
					ans = max(ans, ai+aj) // 找到一个二元组
					cnt2 := 0
					for k := j - 1; k >= 0; k-- {
						ak := a[k]
						if ai%ak != 0 && aj%ak != 0 {
							ans = max(ans, ai+aj+ak) // 找到一个三元组
							if cnt2++; cnt2 == 10 {
								break
							}
						}
					}
					if cnt++; cnt == 10 {
						break
					}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1183F(os.Stdin, os.Stdout) }
