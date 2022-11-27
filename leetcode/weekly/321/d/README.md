#### 提示 1

把比 $k$ 大的数变成 $1$，比 $k$ 小的数变成 $-1$，$k$ 变成 $0$。

#### 提示 2

设 $k$ 的下标为 $\textit{pos}$，$k$ 为子数组的中位数，等价于：

1. 子数组包含下标 $\textit{pos}$；
2. 子数组的元素和等于 $0$ 或 $1$。

#### 提示 3

统计子数组 $\textit{nums}[\textit{pos}..i]$ 中比 $k$ **大**的数的个数，减去比 $k$ **小**的数的个数，记作 $c_i$。

用哈希表 $\textit{cnt}$ 统计 $c_i$ 的个数。

然后对于子数组 $\textit{nums}[i..\textit{pos}]$，统计比 $k$ **小**的数的个数，减去比 $k$ **大**的数的个数，记作 $c_i$。

因此，对于每个 $i$：

- $\textit{cnt}[c_i]$ 就是符合提示 2 的奇数长度子数组的个数；
- $\textit{cnt}[c_i+1]$ 就是符合提示 2 的偶数长度子数组的个数。

累加个数，即为答案。

> 另外一种理解方式是，$k$ 为奇数长度子数组的中位数，等价于左侧小于 + 右侧小于 = 左侧大于 + 右侧大于
>
> 即「左侧小于 - 左侧大于 = 右侧大于 - 右侧小于」

```py [sol1-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        pos = nums.index(k)
        cnt = defaultdict(int)
        cnt[0] = 1  # i=pos 的时候 c 是 0，直接记到 cnt 中，这样下面不是大于就是小于
        c = 0
        for i in range(pos + 1, len(nums)):
            c += 1 if nums[i] > k else -1
            cnt[c] += 1

        ans = cnt[0] + cnt[1]  # i=pos 的时候 c 是 0，直接加到答案中，这样下面不是大于就是小于
        c = 0
        for i in range(pos - 1, -1, -1):
            c += 1 if nums[i] < k else -1
            ans += cnt[c] + cnt[c + 1]
        return ans
```

```java [sol1-Java]
class Solution {
    public int countSubarrays(int[] nums, int k) {
        int pos = 0, n = nums.length;
        while (nums[pos] != k) ++pos;

        var cnt = new HashMap<Integer, Integer>();
        cnt.put(0, 1); // i=pos 的时候 c 是 0，直接记到 cnt 中，这样下面不是大于就是小于
        for (int i = pos + 1, c = 0; i < n; ++i) {
            c += nums[i] > k ? 1 : -1;
            cnt.put(c, cnt.getOrDefault(c, 0) + 1);
        }

        int ans = cnt.get(0) + cnt.getOrDefault(1, 0); // i=pos 的时候 c 是 0，直接加到答案中，这样下面不是大于就是小于
        for (int i = pos - 1, c = 0; i >= 0; --i) {
            c += nums[i] < k ? 1 : -1;
            ans += cnt.getOrDefault(c, 0) + cnt.getOrDefault(c + 1, 0);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int countSubarrays(vector<int> &nums, int k) {
        int pos = find(nums.begin(), nums.end(), k) - nums.begin(), n = nums.size();
        unordered_map<int, int> cnt;
        cnt[0] = 1; // i=pos 的时候 c 是 0，直接记到 cnt 中，这样下面不是大于就是小于
        for (int i = pos + 1, c = 0; i < n; ++i) {
            c += nums[i] > k ? 1 : -1;
            ++cnt[c];
        }

        int ans = cnt[0] + cnt[1]; // i=pos 的时候 c 是 0，直接加到答案中，这样下面不是大于就是小于
        for (int i = pos - 1, c = 0; i >= 0; --i) {
            c += nums[i] < k ? 1 : -1;
            ans += cnt[c] + cnt[c + 1];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countSubarrays(nums []int, k int) int {
	pos := 0
	for nums[pos] != k {
		pos++
	}

	cnt, c := map[int]int{0: 1}, 0 // i=pos 的时候 c 是 0，直接记到 cnt 中，这样下面不是大于就是小于
	for _, x := range nums[pos+1:] {
		if x > k {
			c++
		} else {
			c--
		}
		cnt[c]++
	}

	ans := cnt[0] + cnt[1] // i=pos 的时候 c 是 0，直接加到答案中，这样下面不是大于就是小于
	for i, c := pos-1, 0; i >= 0; i-- {
		if nums[i] < k {
			c++
		} else {
			c--
		}
		ans += cnt[c] + cnt[c+1]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
