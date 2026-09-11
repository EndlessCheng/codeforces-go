## 方法一：暴力枚举

枚举从 $\textit{digits}$ 中选三个数（$A_n^3$ 种选法），分别作为个位数（必须是偶数）、十位数和百位数（不能是 $0$）。把生成的三位数添加到一个哈希集合中。

最后答案为哈希集合的大小。

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

## 方法二：枚举个位数 + 组合数学

偶数的个位数必须是偶数。

枚举个位数为 $d = 0,2,4,6,8$。

消耗一个 $d$ 后，在**剩余数字**中：

- 设有 $\textit{kinds}$ 种数字，那么十位有 $\textit{kinds}$ 种填法。
- 设有 $\textit{nonZeros}$ 种非零数字，那么百位有 $\textit{nonZeros}$ 种填法。
- 设有 $\textit{singles}$ 种恰好出现一次的非零数字，那么十位百位都填这种数字的方案是不合法的，要减去。

所以百位和十位有

$$
\textit{kinds}\cdot \textit{nonZeros} - \textit{singles}
$$

种填法。

```py [sol-Python3]
class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        cnt = Counter(digits)

        kinds = len(cnt)
        non_zeros = kinds - (0 in cnt)
        singles = sum(c == 1 for c in cnt.values()) - (cnt.get(0, 0) == 1)
        ans = 0

        # 枚举个位填偶数 d
        for d, c in cnt.items():
            if d % 2 > 0:
                continue

            # 十位填任意数字
            k = kinds - (c == 1)

            # 百位填任意非零数字
            nz = non_zeros - (d > 0 and c == 1)

            # 恰好出现一次的非零数字，不能同时填入十位和百位
            s = singles
            if d > 0:
                if c == 1:
                    s -= 1
                elif c == 2:
                    s += 1

            ans += k * nz - s

        return ans
```

```java [sol-Java]
class Solution {
    public int totalNumbers(int[] digits) {
        int[] cnt = new int[10];
        for (int d : digits) {
            cnt[d]++;
        }

        int nonZeros = 0;
        int kinds = 0;
        int singles = 0;
        for (int d = 0; d < 10; d++) {
            if (cnt[d] == 0) {
                continue;
            }
            kinds++;
            if (d > 0) {
                nonZeros++;
                if (cnt[d] == 1) {
                    singles++;
                }
            }
        }

        int ans = 0;

        // 枚举个位填偶数 d
        for (int d = 0; d < 10; d += 2) {
            int c = cnt[d];
            if (c == 0) {
                continue;
            }

            // 十位填任意数字
            int k = kinds;
            if (c == 1) {
                k--;
            }

            // 百位填任意非零数字
            int nz = nonZeros;
            if (d > 0 && c == 1) {
                nz--;
            }

            // 恰好出现一次的非零数字，不能同时填入十位和百位
            int s = singles;
            if (d > 0) {
                if (c == 1) {
                    s--;
                } else if (c == 2) {
                    s++; // 个位数填入 d 后，d 恰好出现一次
                }
            }

            ans += k * nz - s;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int totalNumbers(vector<int>& digits) {
        int cnt[10]{};
        for (int d : digits) {
            cnt[d]++;
        }

        int non_zeros = 0, kinds = 0, singles = 0;
        for (int d = 0; d < 10; d++) {
            if (cnt[d] == 0) {
                continue;
            }
            kinds++;
            if (d > 0) {
                non_zeros++;
                singles += cnt[d] == 1;
            }
        }

        int ans = 0;

        // 枚举个位填偶数 d
        for (int d = 0; d < 10; d += 2) {
            int c = cnt[d];
            if (c == 0) {
                continue;
            }

            // 十位填任意数字
            int k = kinds - (c == 1);

            // 百位填任意非零数字
            int nz = non_zeros - (d > 0 && c == 1);

            // 恰好出现一次的非零数字，不能同时填入十位和百位
            int s = singles;
            if (d > 0) {
                if (c == 1) {
                    s--;
                } else if (c == 2) {
                    s++; // 个位数填入 d 后，d 恰好出现一次
                }
            }

            ans += k * nz - s;
        }

        return ans;
    }
};
```

```go [sol-Go]
func totalNumbers(digits []int) (ans int) {
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d]++
	}

	var nonZeros, kinds, singles int
	for d, c := range cnt {
		if c == 0 {
			continue
		}
		kinds++
		if d > 0 {
			nonZeros++
			if c == 1 {
				singles++
			}
		}
	}

	// 枚举个位填偶数 d
	for d := 0; d < 10; d += 2 {
		c := cnt[d]
		if c == 0 {
			continue
		}

		// 十位填任意数字
		k := kinds
		if c == 1 {
			k--
		}

		// 百位填任意非零数字
		nz := nonZeros
		if d > 0 && c == 1 {
			nz--
		}

		// 恰好出现一次的非零数字，不能同时填入十位和百位
		s := singles
		if d > 0 {
			if c == 1 {
				s--
			} else if c == 2 {
				s++ // 个位数填入 d 后，d 恰好出现一次
			}
		}

		ans += k*nz - s
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+D)$，其中 $n$ 是 $\textit{digits}$ 的长度，$D=10$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(D)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
