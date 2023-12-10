[本题视频讲解](https://www.bilibili.com/video/BV1Lj411s7ga/)

考虑如下数组

$$
[3,1,2,1,2,4,4]
$$

题目要求相同数字必须在同一个子数组中，所以两个 $1$ 必须在同一个子数组，两个 $2$ 也必须在同一个子数组。所以 $[1,2,1,2]$ 这一段必须是完整的，不能分割。

把该数组分到无法再分，得到

$$
[3] + [1,2,1,2] + [4,4]
$$

考虑每个 $+$ 号**选或不选**，一共有 $2^2=4$ 种好分割方案，即

$$
\begin{aligned}
&[3] + [1,2,1,2] + [4,4]\\
&[3] + [1,2,1,2,4,4]\\
&[3,1,2,1,2] + [4,4]\\
&[3,1,2,1,2,4,4]
\end{aligned}
$$

### 写法一：合并区间

代码实现时，用一个哈希表/有序集合记录每个元素首次出现的位置和最后一次出现的位置，每个元素就对应着一个不可分割的区间。然后按照 [56. 合并区间](https://leetcode.cn/problems/merge-intervals/) 的做法，把这些区间都合并起来。假设合并后的区间个数为 $m$，那么答案就是

$$
2^{m-1}
$$

记得取模。

注意代码中少统计了最后一段区间，所以直接算的是 $2^m$。

```py [sol-Python3]
class Solution:
    def numberOfGoodPartitions(self, nums: List[int]) -> int:
        ps = {}
        for i, x in enumerate(nums):
            if x in ps:
                ps[x][1] = i  # 更新区间右端点
            else:
                ps[x] = [i, i]

        a = sorted(ps.values(), key=lambda p: p[0])  # 按区间左端点排序

        m = 0
        max_r = a[0][1]
        for left, right in a[1:]:
            if left > max_r:  # 无法合并
                m += 1
            max_r = max(max_r, right)
        return pow(2, m, 1_000_000_007)
```

```java [sol-Java]
class Solution {
    public int numberOfGoodPartitions(int[] nums) {
        Map<Integer, int[]> ps = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (ps.containsKey(x)) {
                ps.get(x)[1] = i; // 更新区间右端点
            } else {
                ps.put(x, new int[]{i, i});
            }
        }

        List<int[]> a = new ArrayList<>(ps.values());
        a.sort((p, q) -> p[0] - q[0]); // 按区间左端点排序

        int ans = 1;
        int maxR = a.get(0)[1];
        for (int i = 1; i < a.size(); i++) {
            int[] interval = a.get(i);
            int left = interval[0];
            int right = interval[1];
            if (left > maxR) { // 无法合并
                ans = ans * 2 % 1_000_000_007;
            }
            maxR = Math.max(maxR, right);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfGoodPartitions(vector<int> &nums) {
        unordered_map<int, pair<int, int>> ps;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            auto it = ps.find(x);
            if (it != ps.end()) {
                it->second.second = i; // 更新区间右端点
            } else {
                ps[x] = {i, i};
            }
        }

        vector<pair<int, int>> a;
        for (auto &[_, p]: ps) {
            a.emplace_back(p);
        }
        sort(a.begin(), a.end(), [](const auto &p, const auto &q) {
            return p.first < q.first; // 按区间左端点排序
        });

        int ans = 1;
        int max_r = a[0].second;
        for (int i = 1; i < a.size(); i++) {
            int left = a[i].first, right = a[i].second;
            if (left > max_r) { // 无法合并
                ans = ans * 2 % 1'000'000'007;
            }
            max_r = max(max_r, right);
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfGoodPartitions(nums []int) int {
	type pair struct{ l, r int }
	ps := map[int]pair{}
	for i, x := range nums {
		if p, ok := ps[x]; ok {
			p.r = i // 更新区间右端点
			ps[x] = p
		} else {
			ps[x] = pair{i, i}
		}
	}

	a := make([]pair, 0, len(ps))
	for _, p := range ps {
		a = append(a, p)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.l - b.l }) // 按区间左端点排序

	ans := 1
	maxR := a[0].r
	for _, p := range a[1:] {
		if p.l > maxR { // 无法合并
			ans = ans * 2 % 1_000_000_007
		}
		maxR = max(maxR, p.r)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 写法二：只记录右端点

1. 遍历数组，用哈希表 $r$ 记录 $\textit{nums}[i]$ 最右边的下标。
2. 再次遍历数组，那么第一个区间就是 $[0, r[\textit{nums}[0]]]$。
3. 如果第二个区间和第一个区间有交集，那么合并区间，维护合并后的区间的右端点 $\textit{maxR}$。
4. 如果第二个区间和第一个区间没有交集，把合并后的区间个数 $m$ 加一。怎么判断没有交集？只要第一个区间的 $\textit{maxR}=i$ 就表示没有交集。
5. 返回 $2^{m-1}$。记得取模。

```py [sol-Python3]
class Solution:
    def numberOfGoodPartitions(self, nums: List[int]) -> int:
        r = {}
        for i, x in enumerate(nums):
            r[x] = i
        m = max_r = 0
        for i, x in enumerate(nums):
            max_r = max(max_r, r[x])
            if max_r == i:  # 区间无法延长
                m += 1
        return pow(2, m - 1, 1_000_000_007)
```

```java [sol-Java]
class Solution {
    public int numberOfGoodPartitions(int[] nums) {
        HashMap<Integer, Integer> r = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            r.put(nums[i], i);
        }
        int ans = 1, max_r = 0;
        for (int i = 0; i < nums.length - 1; i++) { // 少统计最后一段区间
            max_r = Math.max(max_r, r.get(nums[i]));
            if (max_r == i) { // 区间无法延长
                ans = ans * 2 % 1_000_000_007;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfGoodPartitions(vector<int>& nums) {
        unordered_map<int, int> r;
        for (int i = 0; i < nums.size(); i++) {
            r[nums[i]] = i;
        }
        int ans = 1, max_r = 0;
        for (int i = 0; i + 1 < nums.size(); i++) { // 少统计最后一段区间
            max_r = max(max_r, r[nums[i]]);
            if (max_r == i) { // 区间无法延长
                ans = ans * 2 % 1'000'000'007;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfGoodPartitions(nums []int) int {
	r := map[int]int{}
	for i, x := range nums {
		r[x] = i
	}
	ans := 1
	maxR := 0
	for i, x := range nums[:len(nums)-1] { // 少统计最后一段区间
		maxR = max(maxR, r[x])
		if maxR == i { // 区间无法延长
			ans = ans * 2 % 1_000_000_007
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 相似题目

- [56. 合并区间](https://leetcode.cn/problems/merge-intervals/)
- [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)
- [2580. 统计将重叠区间合并成组的方案数](https://leetcode.cn/problems/count-ways-to-group-overlapping-ranges/) 1632
- [2584. 分割数组使乘积互质](https://leetcode.cn/problems/split-the-array-to-make-coprime-products/) 2159
- [2655. 寻找最大长度的未覆盖区间](https://leetcode.cn/problems/find-maximal-uncovered-ranges/)（会员题）
