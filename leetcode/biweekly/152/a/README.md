枚举从 $\textit{digits}$ 中选三个数（$A_n^3$ 种选法），分别作为个位数（必须是偶数）、十位数和百位数（不能是 $0$）。把生成的三位数加到一个哈希集合中。

最后答案就是哈希集合的大小。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1i6Q8YUEtN/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        st = set()
        for i, a in enumerate(digits):  # 个位数
            if a % 2:
                continue
            for j, b in enumerate(digits):  # 十位数
                if j == i:
                    continue
                for k, c in enumerate(digits):  # 百位数
                    if c == 0 or k == i or k == j:
                        continue
                    st.add(c * 100 + b * 10 + a)
        return len(st)
```

```py [sol-Python3 permutations]
class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        st = set()
        for a, b, c in permutations(digits, 3):
            if c and a % 2 == 0:
                st.add(c * 100 + b * 10 + a)
        return len(st)
```

```py [sol-Python3 简化]
class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        return len(set(c * 100 + b * 10 + a
                       for a, b, c in permutations(digits, 3) if c and a % 2 == 0))
```

```java [sol-Java]
class Solution {
    public int totalNumbers(int[] digits) {
        Set<Integer> set = new HashSet<>();
        int n = digits.length;
        for (int i = 0; i < n; i++) { // 个位数
            int a = digits[i];
            if (a % 2 > 0) {
                continue;
            }
            for (int j = 0; j < n; j++) { // 十位数
                if (j == i) {
                    continue;
                }
                for (int k = 0; k < n; k++) { // 百位数
                    int c = digits[k];
                    if (c == 0 || k == i || k == j) {
                        continue;
                    }
                    set.add(c * 100 + digits[j] * 10 + a);
                }
            }
        }
        return set.size();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int totalNumbers(vector<int>& digits) {
        unordered_set<int> st;
        int n = digits.size();
        for (int i = 0; i < n; i++) { // 个位数
            int a = digits[i];
            if (a % 2) {
                continue;
            }
            for (int j = 0; j < n; j++) { // 十位数
                if (j == i) {
                    continue;
                }
                for (int k = 0; k < n; k++) { // 百位数
                    int c = digits[k];
                    if (c == 0 || k == i || k == j) {
                        continue;
                    }
                    st.insert(c * 100 + digits[j] * 10 + a);
                }
            }
        }
        return st.size();
    }
};
```

```go [sol-Go]
func totalNumbers(digits []int) int {
	set := map[int]struct{}{}
	for i, a := range digits { // 个位数
		if a%2 > 0 {
			continue
		}
		for j, b := range digits { // 十位数
			if j == i {
				continue
			}
			for k, c := range digits { // 百位数
				if c == 0 || k == i || k == j {
					continue
				}
				set[c*100+b*10+a] = struct{}{}
			}
		}
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 是 $\textit{digits}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^3)$。

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
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
