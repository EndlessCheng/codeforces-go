**核心思想**：枚举 $j$ 并**暴力交换** $\textit{nums}[j]$ 中的数位。设交换后的数字为 $x$，统计左边有多少个 $\textit{nums}[i]$ 等于 $x$。

为避免重复统计，可以先用一个哈希集合记录交换后的数字，再去遍历哈希集合中的元素 $x$，统计左边有多少个 $\textit{nums}[i]$ 等于 $x$。

但是，如果 $100$ 这样的数在左边，我们枚举的 $1$ 这样的数在右边，就没法找到近似相等的数对。

怎么办？把 $\textit{nums}$ **按照数字长度从小到大排序**（也就是元素值从小到大排序），即可解决此问题。

⚠**注意**：代码在第二次交换时，$p$ 是从 $i+1$ 开始枚举，而不是从 $i$ 开始枚举的，因为在 $i$ 和 $j$ 以及 $i$ 和 $k$ 上的交换，一定等价于在 $i$ 和 $k$ 以及 $j$ 和 $k$ 上的交换。例如 $456$ 先交换 $i=0$ 和 $j=1$，得 $546$，然后交换 $i=0$ 和 $k=2$，得 $645$；我们也可以先交换 $i=0$ 和 $k=2$，得 $654$，再交换 $j=1$ 和 $k=2$，也可以得到 $645$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cMW6ePEwC/) 第四题，欢迎点赞关注！

## 优化前

```py [sol-Python3]
class Solution:
    def countPairs(self, nums: List[int]) -> int:
        nums.sort()
        ans = 0
        cnt = defaultdict(int)
        for x in nums:
            st = {x}  # 不交换
            s = list(str(x))
            m = len(s)
            for i in range(m):
                for j in range(i + 1, m):
                    s[i], s[j] = s[j], s[i]
                    st.add(int(''.join(s)))  # 交换一次
                    for p in range(i + 1, m):
                        for q in range(p + 1, m):
                            s[p], s[q] = s[q], s[p]
                            st.add(int(''.join(s)))  # 交换两次
                            s[p], s[q] = s[q], s[p]
                    s[i], s[j] = s[j], s[i]
            ans += sum(cnt[v] for v in st)
            cnt[x] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countPairs(int[] nums) {
        Arrays.sort(nums);
        int ans = 0;
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            Set<Integer> st = new HashSet<>();
            st.add(x); // 不交换
            char[] s = Integer.toString(x).toCharArray();
            int m = s.length;
            for (int i = 0; i < m; i++) {
                for (int j = i + 1; j < m; j++) {
                    swap(s, i, j);
                    st.add(Integer.parseInt(new String(s))); // 交换一次
                    for (int p = i + 1; p < m; p++) {
                        for (int q = p + 1; q < m; q++) {
                            swap(s, p, q);
                            st.add(Integer.parseInt(new String(s))); // 交换两次
                            swap(s, p, q);
                        }
                    }
                    swap(s, i, j);
                }
            }
            for (int v : st) {
                ans += cnt.getOrDefault(v, 0);
            }
            cnt.merge(x, 1, Integer::sum);
        }
        return ans;
    }

    private void swap(char[] s, int i, int j) {
        char tmp = s[i];
        s[i] = s[j];
        s[j] = tmp;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPairs(vector<int>& nums) {
        ranges::sort(nums);
        int ans = 0;
        unordered_map<int, int> cnt;
        for (int x : nums) {
            unordered_set<int> st = {x}; // 不交换
            string s = to_string(x);
            int m = s.length();
            for (int i = 0; i < m; i++) {
                for (int j = i + 1; j < m; j++) {
                    swap(s[i], s[j]);
                    st.insert(stoi(s)); // 交换一次
                    for (int p = i + 1; p < m; p++) {
                        for (int q = p + 1; q < m; q++) {
                            swap(s[p], s[q]);
                            st.insert(stoi(s)); // 交换两次
                            swap(s[p], s[q]);
                        }
                    }
                    swap(s[i], s[j]);
                }
            }
            for (int v : st) {
                ans += cnt.contains(v) ? cnt[v] : 0;
            }
            cnt[x]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPairs(nums []int) (ans int) {
	slices.Sort(nums)
	cnt := map[int]int{}
	for _, x := range nums {
		set := map[int]struct{}{x: {}} // 不交换
		s := []byte(strconv.Itoa(x))
		m := len(s)
		for i := range s {
			for j := i + 1; j < m; j++ {
				s[i], s[j] = s[j], s[i]
				set[atoi(s)] = struct{}{} // 交换一次
				for p := i + 1; p < m; p++ {
					for q := p + 1; q < m; q++ {
						s[p], s[q] = s[q], s[p]
						set[atoi(s)] = struct{}{} // 交换两次
						s[p], s[q] = s[q], s[p]
					}
				}
				s[i], s[j] = s[j], s[i]
			}
		}
		for x := range set {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return
}

// 手动转 int 快一些
func atoi(s []byte) (res int) {
	for _, b := range s {
		res = res*10 + int(b&15)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log^5 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log^4 U)$。

## 优化

把 $x=\textit{nums}[i]$ 的数位拆开，放入一个数组 $a$ 中。例如 $x=123$ 拆开得到 $a=[3,2,1]$，注意低位在数组左边，这样下标与 $10$ 的幂次可以对应上。

交换 $a[i]$ 和 $a[j]$，相当于 $x$ 增加了

$$
\begin{aligned}
    & (a[j]-a[i])\cdot 10^i + (a[i]-a[j]) \cdot 10^j      \\
={} & (a[j]-a[i])(10^i-10^j)        \\
\end{aligned}
$$

这样就不需要把字符串再转成数字了，预处理 $10^i$ 后，可以 $\mathcal{O}(1)$ 算出交换后的数字。

此外，我们可以在 $a[i]=a[j]$ 时，直接 `continue`，因为交换不改变 $x$ 的值。

```py [sol-Python3]
POW10 = [10 ** i for i in range(7)]

class Solution:
    def countPairs(self, nums: List[int]) -> int:
        nums.sort()
        ans = 0
        cnt = defaultdict(int)
        for x in nums:
            st = {x}  # 不交换
            a = list(map(int, str(x)))[::-1]
            m = len(a)
            for i in range(m):
                for j in range(i + 1, m):
                    if a[i] == a[j]:  # 小优化
                        continue
                    y = x + (a[j] - a[i]) * (POW10[i] - POW10[j])
                    st.add(y)  # 交换一次
                    a[i], a[j] = a[j], a[i]
                    for p in range(i + 1, m):
                        for q in range(p + 1, m):
                            st.add(y + (a[q] - a[p]) * (POW10[p] - POW10[q]))  # 交换两次
                    a[i], a[j] = a[j], a[i]
            ans += sum(cnt[v] for v in st)
            cnt[x] += 1
        return ans
```

```java [sol-Java]
class Solution {
    private static final int[] POW10 = {1, 10, 100, 1000, 10000, 100000, 1000000};

    public int countPairs(int[] nums) {
        Arrays.sort(nums);
        int ans = 0;
        Map<Integer, Integer> cnt = new HashMap<>();
        int[] a = new int[7];
        for (int x : nums) {
            Set<Integer> st = new HashSet<>();
            st.add(x); // 不交换
            int m = 0;
            for (int v = x; v > 0; v /= 10) {
                a[m++] = v % 10;
            }
            for (int i = 0; i < m; i++) {
                for (int j = i + 1; j < m; j++) {
                    if (a[i] == a[j]) { // 小优化
                        continue;
                    }
                    int y = x + (a[j] - a[i]) * (POW10[i] - POW10[j]);
                    st.add(y); // 交换一次
                    swap(a, i, j);
                    for (int p = i + 1; p < m; p++) {
                        for (int q = p + 1; q < m; q++) {
                            st.add(y + (a[q] - a[p]) * (POW10[p] - POW10[q])); // 交换两次
                        }
                    }
                    swap(a, i, j);
                }
            }
            for (int v : st) {
                ans += cnt.getOrDefault(v, 0);
            }
            cnt.merge(x, 1, Integer::sum);
        }
        return ans;
    }

    private void swap(int[] a, int i, int j) {
        int tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    private static final int[] POW10 = {1, 10, 100, 1000, 10000, 100000, 1000000};

    public int countPairs(int[] nums) {
        Arrays.sort(nums);
        int ans = 0;
        Map<Integer, Integer> cnt = new HashMap<>();
        int[] a = new int[7];
        for (int x : nums) {
            Set<Integer> vis = new HashSet<>();
            vis.add(x); // 不交换
            ans += cnt.getOrDefault(x, 0);
            int m = 0;
            for (int v = x; v > 0; v /= 10) {
                a[m++] = v % 10;
            }
            for (int i = 0; i < m; i++) {
                for (int j = i + 1; j < m; j++) {
                    if (a[i] == a[j]) { // 小优化
                        continue;
                    }
                    int y = x + (a[j] - a[i]) * (POW10[i] - POW10[j]);
                    if (vis.add(y)) {
                        ans += cnt.getOrDefault(y, 0); // 交换一次
                    }
                    swap(a, i, j);
                    for (int p = i + 1; p < m; p++) {
                        for (int q = p + 1; q < m; q++) {
                            int z = y + (a[q] - a[p]) * (POW10[p] - POW10[q]);
                            if (vis.add(z)) {
                                ans += cnt.getOrDefault(z, 0); // 交换两次
                            }
                        }
                    }
                    swap(a, i, j);
                }
            }
            cnt.merge(x, 1, Integer::sum);
        }
        return ans;
    }

    private void swap(int[] a, int i, int j) {
        int tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }
}
```

```cpp [sol-C++]
const int POW10[7] = {1, 10, 100, 1000, 10000, 100000, 1000000};

class Solution {
public:
    int countPairs(vector<int>& nums) {
        ranges::sort(nums);
        int ans = 0, a[7];
        unordered_map<int, int> cnt;
        for (int x : nums) {
            unordered_set<int> st = {x}; // 不交换
            int m = 0;
            for (int v = x; v; v /= 10) {
                a[m++] = v % 10;
            }
            for (int i = 0; i < m; i++) {
                for (int j = i + 1; j < m; j++) {
                    if (a[i] == a[j]) { // 小优化
                        continue;
                    }
                    int y = x + (a[j] - a[i]) * (POW10[i] - POW10[j]);
                    st.insert(y); // 交换一次
                    swap(a[i], a[j]);
                    for (int p = i + 1; p < m; p++) {
                        for (int q = p + 1; q < m; q++) {
                            st.insert(y + (a[q] - a[p]) * (POW10[p] - POW10[q])); // 交换两次
                        }
                    }
                    swap(a[i], a[j]);
                }
            }
            for (int v : st) {
                ans += cnt.contains(v) ? cnt[v] : 0;
            }
            cnt[x]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
var pow10 = [...]int{1, 10, 100, 1000, 10000, 100000, 1000000}

func countPairs(nums []int) int {
	slices.Sort(nums)
	ans := 0
	cnt := make(map[int]int)
	a := [len(pow10)]int{}
	for _, x := range nums {
		st := map[int]struct{}{x: {}} // 不交换
		m := 0
		for v := x; v > 0; v /= 10 {
			a[m] = v % 10
			m++
		}
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				if a[i] == a[j] { // 小优化
					continue
				}
				y := x + (a[j]-a[i])*(pow10[i]-pow10[j])
				st[y] = struct{}{} // 交换一次
				a[i], a[j] = a[j], a[i]
				for p := i + 1; p < m; p++ {
					for q := p + 1; q < m; q++ {
						st[y+(a[q]-a[p])*(pow10[p]-pow10[q])] = struct{}{} // 交换两次
					}
				}
				a[i], a[j] = a[j], a[i]
			}
		}
		for x := range st {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log^4 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log^4 U)$。

更多相似题目，见下面数据结构题单中的「**常用技巧**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
