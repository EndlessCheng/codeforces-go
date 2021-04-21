package main

import (
	. "fmt"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func ioq14() {
	var T, n, sccID int
	var ok bool
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		a := make([]int, n) // 节点编号
		for i := range a {
			a[i] = i
		}
		sort.Slice(a, func(i, j int) (d bool) { Println(1, a[j], a[i]); Scan(&d); return })
		// 竞赛图的 SCC 是一条链
		// 排序后，a[0] 到 a[n-1] 即为从链尾到链首的顺序
		// 从链尾向链首遍历，给各个 SCC 标上 sccID
		// 每个 SCC 内部双指针扫描，只要左指针能访问到一个右指针及其前面的点就说明右指针在 SCC 内
		g := make([]int, n)
		for i := 0; i < n; {
			j := i
			sccID++
			g[a[i]] = sccID
			i++
			for ; j < i; j++ {
				for i < n {
					Print(2, a[j], n-i)
					for _, v := range a[i:] {
						Print(" ", v)
					}
					Println()
					if Scan(&ok); !ok {
						break
					}
					g[a[i]] = sccID
					i++
				}
			}
		}
		Println(3)
		for _, v := range g {
			s := make([]byte, n)
			for j, w := range g {
				if v < w {
					s[j] = '0'
				} else {
					s[j] = '1'
				}
			}
			Println(string(s))
		}
		Scan(&n)
	}
}

//func main() { ioq14() }
