为方便描述，下文把 $\textit{word}$ 简称为 $s$，把 $\textit{numFriends}$ 简称为 $k$。

## 题意

把 $s$ 分割为 $k$ 个非空子串，返回其中字典序最大的子串。

## 方法一：枚举子串左端点

如果固定子串的左端点，那么**子串越长，字典序越大**。

所以核心思路是：枚举子串的左端点，计算最大子串。

单个子串的长度不能超过多少？

由于其余 $k-1$ 个子串必须是非空的，取长度为 $1$，其余子串的长度之和**至少**为 $k-1$。

所以单个子串的长度**至多**为 $n-(k-1)$。

注意特判 $k=1$ 的情况，此时无法分割（子串左端点只能是 $0$），所以答案是 $s$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV13f68YjE7o/?t=5m16s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def answerString(self, s: str, k: int) -> str:
        if k == 1:
            return s
        n = len(s)
        return max(s[i: i + n - k + 1] for i in range(n))
```

```java [sol-Java]
class Solution {
    public String answerString(String s, int k) {
        if (k == 1) {
            return s;
        }
        int n = s.length();
        String ans = "";
        for (int i = 0; i < n; i++) {
            String sub = s.substring(i, Math.min(i + n - k + 1, n));
            if (sub.compareTo(ans) > 0) {
                ans = sub;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string answerString(string s, int k) {
        if (k == 1) {
            return s;
        }
        int n = s.length();
        string ans;
        for (int i = 0; i < n; i++) {
            ans = max(ans, s.substr(i, n - max(k - 1, i)));
        }
        return ans;
    }
};
```

```go [sol-Go]
func answerString(s string, k int) (ans string) {
	if k == 1 {
		return s
	}
	n := len(s)
	for i := range n {
		ans = max(ans, s[i:min(i+n-k+1, n)])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n-k))$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n-k)$ 或 $\mathcal{O}(1)$。Go 的切片不会发生拷贝，只需要 $\mathcal{O}(1)$ 额外空间。

## 方法二：计算字典序最大的后缀

也可以先把字典序最大的后缀算出来，然后取它的长度至多为 $n-k+1$ 的前缀，作为答案。

正确性证明。在比较字典序大小的过程中，如果当前后缀比前面的某个后缀的字典序更大，那么：

- 如果两个后缀的首个不同字母的位置小于 $n-k+1$，那么应当更新答案的最大值。
- 如果两个后缀的首个不同字母的位置不小于 $n-k+1$，那么这两个后缀（只看长为 $n-k+1$ 的前缀）是一样的，即使更新答案，也不影响最终结果。

如何计算字典序最大的后缀，见 [1163. 按字典序排在最后的子串](https://leetcode.cn/problems/last-substring-in-lexicographical-order/)。

```py [sol-Python3]
class Solution:
    def answerString(self, s: str, numFriends: int) -> str:
        if numFriends == 1:
            return s
        n = len(s)
        i, j = 0, 1
        while j < n:
            k = 0
            while j + k < n and s[i + k] == s[j + k]:
                k += 1
            if j + k < n and s[i + k] < s[j + k]:
                i, j = j, max(j + 1, i + k + 1)
            else:
                j += k + 1
        return s[i: i + n - numFriends + 1]
```

```java [sol-Java]
class Solution {
    public String answerString(String s, int numFriends) {
        if (numFriends == 1) {
            return s;
        }
        int n = s.length();
        int i = 0;
        int j = 1;
        while (j < n) {
            int k = 0;
            while (j + k < n && s.charAt(i + k) == s.charAt(j + k)) {
                k++;
            }
            if (j + k < n && s.charAt(i + k) < s.charAt(j + k)) {
                int t = i;
                i = j;
                j = Math.max(j + 1, t + k + 1);
            } else {
                j += k + 1;
            }
        }
        return s.substring(i, Math.min(i + n - numFriends + 1, n));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string answerString(string s, int k) {
        if (k == 1) {
            return s;
        }
        int n = s.length();
        int i = 0, j = 1;
        while (j < n) {
            int k = 0;
            while (j + k < n && s[i + k] == s[j + k]) {
                k++;
            }
            if (j + k < n && s[i + k] < s[j + k]) {
                int t = i;
                i = j;
                j = max(j + 1, t + k + 1);
            } else {
                j += k + 1;
            }
        }
        return s.substr(i, n - max(k - 1, i));
    }
};
```

```go [sol-Go]
func answerString(s string, k int) string {
	if k == 1 {
		return s
	}
	n := len(s)
	i, j := 0, 1
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			k++
		}
		if j+k < n && s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j += k + 1
		}
	}
	return s[i:min(i+n-k+1, n)]
}
```

```go [sol-Go 后缀数组]
func answerString(s string, k int) string {
	if k == 1 {
		return s
	}
	sa := (*struct{_[]byte;sa[]int32})(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	n := len(s)
	i := int(sa[n-1])
	return s[i:min(i+n-k+1, n)]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

更多相似题目，见下面贪心题单中的「**§3.1 字典序最小/最大**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
