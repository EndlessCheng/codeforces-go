下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

暴力枚举 $a,b,c$ 的全排列 $a',b',c'$，然后合并 $a',b'$ 得到最短字符串 $x$，再合并 $x,c'$ 得到最短字符串 $s$。

可能有的同学觉得，$a',b'$ 合并成一个较长的字符串会不会更优？

你可以这样理解：在没有完全包含的前提下，我们相当于在 $b'$ 的左边添加一些字母，右边添加一些字母，得到一个最短字符串 $s$。

可以先特判一下完全包含的情况，对于没有完全包含的情况，必然是合并得越短越好。

取最短且字典序最小的 $s$ 作为答案。

```py [sol-Python3]
class Solution:
    def minimumString(self, a: str, b: str, c: str) -> str:
        def merge(s: str, t: str) -> str:
            # 先特判完全包含的情况
            if t in s: return s
            if s in t: return t
            for i in range(min(len(s), len(t)), 0, -1):
                # 枚举：s 的后 i 个字母和 t 的前 i 个字母是一样的
                if s[-i:] == t[:i]:
                    return s + t[i:]
            return s + t
        return min((merge(merge(a, b), c) for a, b, c in permutations((a, b, c))), key=lambda s: (len(s), s))
```

```go [sol-Go]
func merge(s, t string) string {
	// 先特判完全包含的情况
	if strings.Contains(s, t) {
		return s
	}
	if strings.Contains(t, s) {
		return t
	}
	for i := min(len(s), len(t)); ; i-- {
		// 枚举：s 的后 i 个字母和 t 的前 i 个字母是一样的
		if s[len(s)-i:] == t[:i] {
			return s + t[i:]
		}
	}
}

func minimumString(a, b, c string) (ans string) {
	arr := []string{a, b, c}
	// 枚举 arr 的全排列
	for _, p := range [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}} {
		s := merge(merge(arr[p[0]], arr[p[1]]), arr[p[2]])
		if ans == "" || len(s) < len(ans) || len(s) == len(ans) && s < ans {
			ans = s
		}
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $a,b,c$ 的长度的最大值。
- 空间复杂度：$\mathcal{O}(n)$。
