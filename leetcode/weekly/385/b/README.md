## 方法一：用字符串处理

把 $\textit{arr}_1$ 的所有前缀丢到一个哈希集合中，然后遍历 $\textit{arr}_2$ 的所有前缀，统计在哈希集合中的最长长度。

```py [sol-Python3]
class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        st = set()
        for s in map(str, arr1):
            for i in range(1, len(s) + 1):  # 枚举 s 的前缀长度
                st.add(s[:i])

        ans = 0
        for s in map(str, arr2):
            for i in range(1, len(s) + 1):  # 枚举 s 的前缀长度
                if s[:i] not in st:
                    break
                ans = max(ans, i)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        Set<String> st = new HashSet<>();
        for (int x : arr1) {
            String s = Integer.toString(x);
            for (int i = 1; i <= s.length(); i++) { // 枚举 s 的前缀长度
                st.add(s.substring(0, i));
            }
        }

        int ans = 0;
        for (int x : arr2) {
            String s = Integer.toString(x);
            for (int i = 1; i <= s.length(); i++) { // 枚举 s 的前缀长度
                if (!st.contains(s.substring(0, i))) {
                    break;
                }
                ans = Math.max(ans, i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestCommonPrefix(vector<int>& arr1, vector<int>& arr2) {
        unordered_set<string> st;
        for (int x : arr1) {
            string s = to_string(x);
            for (int i = 1; i <= s.size(); i++) { // 枚举 s 的前缀长度
                st.insert(s.substr(0, i));
            }
        }

        int ans = 0;
        for (int x : arr2) {
            string s = to_string(x);
            for (int i = 1; i <= s.size(); i++) { // 枚举 s 的前缀长度
                if (!st.contains(s.substr(0, i))) {
                    break;
                }
                ans = max(ans, i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestCommonPrefix(arr1, arr2 []int) (ans int) {
	has := map[string]bool{}
	for _, x := range arr1 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s); i++ { // 枚举 s 的前缀长度
			has[s[:i]] = true
		}
	}

	for _, x := range arr2 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s) && has[s[:i]]; i++ { // 枚举 s 的前缀长度
			ans = max(ans, i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log^2 U)$，其中 $n$ 是 $\textit{arr}_1$ 的长度，$m$ 是 $\textit{arr}_2$ 的长度，$U$ 是数组元素的最大值。值为 $x$ 的数的十进制长度为 $\mathcal{O}(\log x)$。
- 空间复杂度：$\mathcal{O}(n\log^2 U)$。

## 方法二：用整数处理

不断地把元素 $x$ 除以 $10$（下取整）直到 $0$，例如 $123\to 12\to 1\to 0$，这个过程中的数即为 $x$ 的前缀。

代码实现时，不需要计算长度，而是取数值的最大值，因为数值越大长度越长。

```py [sol-Python3]
class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        st = set()
        for x in arr1:
            while x and x not in st:  # 如果 x 在 st 中，那么剩余前缀也在 st 中
                st.add(x)
                x //= 10

        mx = 0
        for x in arr2:
            while x and x not in st:
                x //= 10
            mx = max(mx, x)
        return len(str(mx)) if mx else 0
```

```java [sol-Java]
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        Set<Integer> st = new HashSet<>();
        for (int x : arr1) {
            while (x > 0 && st.add(x)) { // 如果 x 在 st 中，那么剩余前缀也在 st 中
                x /= 10;
            }
        }

        int mx = 0;
        for (int x : arr2) {
            while (x > 0 && !st.contains(x)) {
                x /= 10;
            }
            mx = Math.max(mx, x);
        }
        return mx > 0 ? Integer.toString(mx).length() : 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestCommonPrefix(vector<int>& arr1, vector<int>& arr2) {
        unordered_set<int> st;
        for (int x : arr1) {
            while (x > 0 && st.insert(x).second) { // 如果 x 在 st 中，那么剩余前缀也在 st 中
                x /= 10;
            }
        }

        int mx = 0;
        for (int x : arr2) {
            while (x > 0 && !st.contains(x)) {
                x /= 10;
            }
            mx = max(mx, x);
        }
        return mx ? to_string(mx).length() : 0;
    }
};
```

```go [sol-Go]
func longestCommonPrefix(arr1, arr2 []int) int {
	has := map[int]bool{}
	for _, x := range arr1 {
		for x > 0 && !has[x] { // 如果 x 在 st 中，那么剩余前缀也在 st 中
			has[x] = true
			x /= 10
		}
	}

	mx := 0
	for _, x := range arr2 {
		for x > 0 && !has[x] {
			x /= 10
		}
		mx = max(mx, x)
	}

	if mx == 0 {
		return 0
	}
	return len(strconv.Itoa(mx))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log U)$，其中 $n$ 是 $\textit{arr}_1$ 的长度，$m$ 是 $\textit{arr}_2$ 的长度，$U$ 是数组元素的最大值。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
