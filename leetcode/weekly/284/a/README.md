## 方法一：记录 key 最后一次出现的位置

### 分析

问题相当于，对于每个 $i$，判断在下标区间 $[\max(i-k,0),\min(i+k,n-1)]$ 中，是否有元素等于 $\textit{key}$。

这是一个长度至多为 $2k+1$ 的**滑动窗口**。

可以统计窗口中的 $\textit{key}$ 的出现次数，如果出现次数大于 $0$，就把 $i$ 加入答案。

更简单的做法是，记录 $\textit{key}$ 最后一次出现的位置 $\textit{last}$，只要 $\textit{last}$ 在窗口中，就把 $i$ 加入答案。

### 算法

第一个窗口是 $[0,k]$。

在计算第一个窗口之前，先求出 $[0,k-1]$ 中 $\textit{key}$ 的最后一次出现位置 $\textit{last}$。

然后遍历 $\textit{nums}$，处理每个窗口。如果窗口右端点 $i+k<n$ 且 $\textit{nums}[i+k]=\textit{key}$，那么更新 $\textit{last}=i+k$。

如果 $\textit{last}\ge i-k$，则 $\textit{last}$ 在窗口中，把 $i$ 加入答案。

> **注**：这个做法每个 $\textit{nums}[i]$ 至多遍历一次。

```py [sol-Python3]
class Solution:
    def findKDistantIndices(self, nums: List[int], key: int, k: int) -> List[int]:
        last = -inf
        for i in range(k - 1, -1, -1):
            if nums[i] == key:
                last = i
                break

        ans = []
        n = len(nums)
        for i in range(n):
            if i + k < n and nums[i + k] == key:
                last = i + k
            if last >= i - k:  # key 在窗口中
                ans.append(i)
        return ans
```

```java [sol-Java]
class Solution {
    public List<Integer> findKDistantIndices(int[] nums, int key, int k) {
        int last = -k - 1; // 保证 key 不存在时 last < i-k
        for (int i = k - 1; i >= 0; i--) {
            if (nums[i] == key) {
                last = i;
                break;
            }
        }

        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            if (i + k < nums.length && nums[i + k] == key) {
                last = i + k;
            }
            if (last >= i - k) { // key 在窗口中
                ans.add(i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findKDistantIndices(vector<int>& nums, int key, int k) {
        int last = -k - 1; // 保证 key 不存在时 last < i-k
        for (int i = k - 1; i >= 0; i--) {
            if (nums[i] == key) {
                last = i;
                break;
            }
        }

        vector<int> ans;
        for (int i = 0; i < nums.size(); i++) {
            if (i + k < nums.size() && nums[i + k] == key) {
                last = i + k;
            }
            if (last >= i - k) { // key 在窗口中
                ans.push_back(i);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int* findKDistantIndices(int* nums, int numsSize, int key, int k, int* returnSize) {
    int last = -k - 1; // 保证 key 不存在时 last < i-k
    for (int i = k - 1; i >= 0; i--) {
        if (nums[i] == key) {
            last = i;
            break;
        }
    }

    int* ans = malloc(sizeof(int) * numsSize);
    *returnSize = 0;
    for (int i = 0; i < numsSize; i++) {
        if (i + k < numsSize && nums[i + k] == key) {
            last = i + k;
        }
        if (last >= i - k) { // key 在窗口中
            ans[(*returnSize)++] = i;
        }
    }
    return ans;
}
```

```go [sol-Go]
func findKDistantIndices(nums []int, key, k int) (ans []int) {
	last := -k - 1 // 保证 key 不存在时 last < i-k
	for i := k - 1; i >= 0; i-- {
		if nums[i] == key {
			last = i
			break
		}
	}

	for i := range nums {
		if i+k < len(nums) && nums[i+k] == key {
			last = i + k
		}
		if last >= i-k { // key 在窗口中
			ans = append(ans, i)
		}
	}
	return
}
```

```js [sol-JavaScript]
var findKDistantIndices = function(nums, key, k) {
    let last = -k - 1; // 保证 key 不存在时 last < i-k
    for (let i = k - 1; i >= 0; i--) {
        if (nums[i] === key) {
            last = i;
            break;
        }
    }

    const ans = [];
    for (let i = 0; i < nums.length; i++) {
        if (i + k < nums.length && nums[i + k] === key) {
            last = i + k;
        }
        if (last >= i - k) { // key 在窗口中
            ans.push(i);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_k_distant_indices(nums: Vec<i32>, key: i32, k: i32) -> Vec<i32> {
        let n = nums.len();
        let k = k as usize;

        let mut last = n;
        for i in (0..k).rev() {
            if nums[i] == key {
                last = i;
                break;
            }
        }

        let mut ans = vec![];
        for i in 0..n {
            if i + k < n && nums[i + k] == key {
                last = i + k;
            }
            if last < n && last + k >= i { // last >= i - k
                ans.push(i as i32);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 方法二：同向双指针

方法一是一个一个判断 $i$ 是否在答案中。枚举 $\textit{i}$ 去看 $\textit{key}$。

也可以枚举 $\textit{key}$ 的位置，去把邻近的 $i$ 加入答案。 

具体来说，遍历 $\textit{nums}$，如果 $\textit{nums}[i] = \textit{key}$，那么 $[\max(i-k,0),\min(i+k,n-1)]$ 中的下标可以在答案中。

由于随着 $i$ 的变大，窗口在右移，我们可以用另一个指针 $j$ 维护未加入答案的最小下标。

```py [sol-Python3]
class Solution:
    def findKDistantIndices(self, nums: List[int], key: int, k: int) -> List[int]:
        ans = []
        j = 0
        for i, x in enumerate(nums):
            if x != key:
                continue
            j = max(j, i - k)  # j 至少是 i-k
            while j <= min(i + k, len(nums) - 1):  # j 至多是 i+k，但不能越界
                ans.append(j)  # 另一种写法见【Python3 写法二】
                j += 1
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def findKDistantIndices(self, nums: List[int], key: int, k: int) -> List[int]:
        ans = []
        j = 0
        for i, x in enumerate(nums):
            if x == key:
                # 至少是 i-k，至多是 i+k，但不能越界
                ans.extend(range(max(i - k, j), min(i + k + 1, len(nums))))
                j = i + k + 1
        return ans
```

```java [sol-Java]
class Solution {
    public List<Integer> findKDistantIndices(int[] nums, int key, int k) {
        List<Integer> ans = new ArrayList<>();
        int j = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != key) {
                continue;
            }
            j = Math.max(j, i - k); // j 至少是 i-k
            while (j <= Math.min(i + k, nums.length - 1)) { // j 至多是 i+k，但不能越界
                ans.add(j++);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findKDistantIndices(vector<int>& nums, int key, int k) {
        vector<int> ans;
        int n = nums.size(), j = 0;
        for (int i = 0; i < n; i++) {
            if (nums[i] != key) {
                continue;
            }
            j = max(j, i - k); // j 至少是 i-k
            while (j <= min(i + k, n - 1)) { // j 至多是 i+k，但不能越界
                ans.push_back(j++);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int* findKDistantIndices(int* nums, int numsSize, int key, int k, int* returnSize) {
    int* ans = malloc(sizeof(int) * numsSize);
    *returnSize = 0;
    int j = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] != key) {
            continue;
        }
        j = MAX(j, i - k); // j 至少是 i-k
        while (j <= MIN(i + k, numsSize - 1)) { // j 至多是 i+k，但不能越界
            ans[(*returnSize)++] = j++;
        }
    }
    return ans;
}
```

```go [sol-Go]
func findKDistantIndices(nums []int, key, k int) (ans []int) {
	j := 0
	for i, x := range nums {
		if x != key {
			continue
		}
		j = max(j, i-k) // j 至少是 i-k
		for j <= min(i+k, len(nums)-1) { // j 至多是 i+k，但不能越界
			ans = append(ans, j)
			j++
		}
	}
	return
}
```

```js [sol-JavaScript]
var findKDistantIndices = function(nums, key, k) {
    const ans = [];
    let j = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] !== key) {
            continue;
        }
        j = Math.max(j, i - k); // j 至少是 i-k
        while (j <= Math.min(i + k, nums.length - 1)) { // j 至多是 i+k，但不能越界
            ans.push(j++);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_k_distant_indices(nums: Vec<i32>, key: i32, k: i32) -> Vec<i32> {
        let n = nums.len();
        let k = k as usize;
        let mut ans = vec![];
        let mut j = 0;
        for (i, x) in nums.into_iter().enumerate() {
            if x != key {
                continue;
            }
            j = j.max(i.saturating_sub(k)); // j 至少是 i-k
            while j <= (i + k).min(n - 1) { // j 至多是 i+k，但不能越界
                ans.push(j as i32);
                j += 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然写了个二重循环，但是内层循环中对 $j$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
