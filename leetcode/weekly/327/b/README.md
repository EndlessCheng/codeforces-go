[视频讲解](https://www.bilibili.com/video/BV1KG4y1j73o/) 第二题。

用一个最大堆模拟。循环 $k$ 次，每次循环把堆顶加入答案，然后把堆顶 `top` 更新为 `ceil(top/3)`。

原地堆化（heapify）可以做到 $\mathcal{O}(1)$ 的空间复杂度。

```py [sol-Python3]
class Solution:
    def maxKelements(self, nums: List[int], k: int) -> int:
        for i in range(len(nums)):
            nums[i] = -nums[i]  # 最大堆
        heapify(nums)  # 原地堆化
        ans = 0
        for _ in range(k):
            ans -= heapreplace(nums, nums[0] // 3)
        return ans
```

```cpp [sol-C++]
class Solution {
public:
    long long maxKelements(vector<int> &nums, int k) {
        make_heap(nums.begin(), nums.end()); // 原地堆化（最大堆）
        long long ans = 0;
        while (k--) {
            pop_heap(nums.begin(), nums.end()); // 把堆顶移到末尾
            ans += nums.back();
            nums.back() = (nums.back() + 2) / 3;
            push_heap(nums.begin(), nums.end()); // 把末尾元素入堆
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxKelements(nums []int, k int) (ans int64) {
	h := hp{nums}
	heap.Init(&h) // 原地堆化
	for ; k > 0; k-- {
		ans += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
```

```rust [sol-Rust]
use std::collections::BinaryHeap;

impl Solution {
    pub fn max_kelements(nums: Vec<i32>, k: i32) -> i64 {
        let mut h = BinaryHeap::from(nums); // 原地堆化（最大堆）
        let mut ans = 0i64;
        for _ in 0..k {
            let mx = h.pop().unwrap();
            ans += mx as i64;
            h.push((mx + 2) / 3);
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + k\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。把 $\textit{nums}$ 堆化需要 $\mathcal{O}(n)$ 时间。
- 空间复杂度：$\mathcal{O}(1)$。

欢迎关注 [B站@灵茶山艾府](https://b23.tv/JMcHRRp)

[往期题解精选（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
