本题和周赛第四题类似，但至多交换一次，请看 [我的题解](https://leetcode.cn/problems/count-almost-equal-pairs-ii/solutions/2892072/pai-xu-mei-ju-you-wei-hu-zuo-pythonjavac-vbg0/)。

把第二次交换的代码移除即可得到本题的代码。

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

- 时间复杂度：$\mathcal{O}(n\log n + n\log^3 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log^2 U)$。

## 优化后

原理请看请看 [周赛第四题题解](https://leetcode.cn/problems/count-almost-equal-pairs-ii/solutions/2892072/pai-xu-mei-ju-you-wei-hu-zuo-pythonjavac-vbg0/)。

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
                    st.add(x + (a[j] - a[i]) * (POW10[i] - POW10[j]))  # 交换一次
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
                    st.add(x + (a[j] - a[i]) * (POW10[i] - POW10[j])); // 交换一次
                }
            }
            for (int v : st) {
                ans += cnt.getOrDefault(v, 0);
            }
            cnt.merge(x, 1, Integer::sum);
        }
        return ans;
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
                }
            }
            cnt.merge(x, 1, Integer::sum);
        }
        return ans;
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
                    st.insert(x + (a[j] - a[i]) * (POW10[i] - POW10[j])); // 交换一次
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
				st[x+(a[j]-a[i])*(pow10[i]-pow10[j])] = struct{}{} // 交换一次
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

- 时间复杂度：$\mathcal{O}(n\log n + n\log^2 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log^2 U)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
