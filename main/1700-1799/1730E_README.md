### 算法框架

枚举 $v=a_i$ 及其因子 $d$（最多 $240$ 个因子），计算子数组左边界的范围，和右边界的范围，使左右边界之内的数，最大值恰好是 $v$ 且最小值恰好是 $d$。

### 细节

看上去思路比较简单，但是实现起来还是有些技巧的。

1. 预处理每个数的因子列表，这可以反向枚举因子及其倍数得到。
2. 用 [单调栈](https://www.bilibili.com/video/BV1VN411J7S7/) 预处理：
   - $a_i$ 左边最近更大或相等元素的下标 $\textit{leftHi}_i$。如果不存在则为 $-1$。
   - $a_i$ 右边最近更大元素的下标 $\textit{rightHi}_i$。如果不存在则为 $n$。
   - $a_i$ 左边最近更小元素的下标 $\textit{leftLo}_i$。如果不存在则为 $-1$。
   - $a_i$ 右边最近更小元素的下标 $\textit{rightLo}_i$。如果不存在则为 $n$。
3. 预处理 $a$ 中相同元素的下标列表 $\textit{pos}$。其中 $\textit{pos}_v$ 表示 $v$ 在 $a$ 中的下标列表（下标从小到大）。
4. 枚举 $v=a_i$ 及其因子 $d$。设 $l=\textit{leftHi}_i,\ r=\textit{rightHi}_i$。
5. 第一种情况：如果 $v$ 左侧最近 $d$ 存在，设其下标为 $j$，那么子数组左端点范围为 $(\max(\textit{leftLo}_j, l), j]$，右端点范围为 $[i, \min(\textit{rightLo}_j, r))$，两个范围长度的乘积加入答案中。注意前提是 $j > l$ 且 $\textit{rightLo}_j > i$。然后更新 $l$ 为 $\max(l,j)$，防止情况二重复统计。
6. 第二种情况：如果 $v$ 右侧最近 $d$ 存在，设其下标为 $j$，那么子数组左端点范围为 $(\max(\textit{leftLo}_j, l), i]$，右端点范围为 $[j, \min(\textit{rightLo}_j, r))$，两个范围长度的乘积加入答案中。注意前提是 $j < r$ 且 $\textit{leftLo}_j < i$。

怎么找 $v$ 左右最近的 $d$ 的下标？在遍历 $a$ 的同时维护 $\textit{pos}$ 列表，对于在 $a_i$ 左侧的相同数字，我们只保留最大的小于 $i$ 的下标。

下面的实现单独判断了 $d=v$ 的情况，此时子数组左端点为 $i$，右端点范围为 $[i, \min(\textit{rightLo}_i, r))$。

```go
package main
import ("bufio";."fmt";"os")
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }

func main() {
	in := bufio.NewReader(os.Stdin)
	const mx = 1000001
	divisors := [mx][]uint32{}
	for i := uint32(1); i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	var T, n int
	pos := [mx][]int{}
	for Fscan(in, &T); T > 0; T-- {
		for i := range pos {
			pos[i] = pos[i][:0]
		}
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			pos[a[i]] = append(pos[a[i]], i)
		}

		leftHi := make([]int, n)  // >= a[i]
		rightHi := make([]int, n) // > a[i]
		leftLo := make([]int, n)  // < a[i]
		rightLo := make([]int, n) // < a[i]
		s := []int{-1} // 哨兵
		t := []int{-1}
		for i, v := range a {
			for len(s) > 1 && v > a[s[len(s)-1]] {
				rightHi[s[len(s)-1]] = i
				s = s[:len(s)-1]
			}
			leftHi[i] = s[len(s)-1]
			s = append(s, i)

			for len(t) > 1 && v <= a[t[len(t)-1]] {
				t = t[:len(t)-1]
			}
			leftLo[i] = t[len(t)-1]
			t = append(t, i)
		}
		for _, i := range s[1:] {
			rightHi[i] = n
		}

		t = append(t[:0], n)
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for len(t) > 1 && v <= a[t[len(t)-1]] {
				t = t[:len(t)-1]
			}
			rightLo[i] = t[len(t)-1]
			t = append(t, i)
		}

		ans := 0
		for i, v := range a { // 最大值为 v
			r := rightHi[i]
			ans += min(rightLo[i], r) - i // 最小值为 v
			for _, d := range divisors[v] { // 最小值为 d
				ps := pos[d]
				l := leftHi[i]
				if len(ps) > 0 && ps[0] < i {
					j := ps[0] // v 左侧最近 d 的下标
					ps = ps[1:]
					if j > l && rightLo[j] > i {
						ans += (j - max(leftLo[j], l)) * (min(rightLo[j], r) - i)
					}
					l = max(l, j) // 避免重复统计
				}
				if len(ps) > 0 {
					j := ps[0] // v 右侧最近 d 的下标
					if j < r && leftLo[j] < i {
						ans += (i - max(leftLo[j], l)) * (min(rightLo[j], r) - j)
					}
				}
			}
			// v 左侧每个数只保留其最右的下标
			if len(pos[v]) > 1 && pos[v][1] == i {
				pos[v] = pos[v][1:]
			}
		}
		Println(ans)
	}
}
```

- 时间复杂度：预处理 $\mathcal{O}(U\log U)$，其中 $U=10^6$。每组数据 $\mathcal{O}(nD)$，其中 $D=240$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
