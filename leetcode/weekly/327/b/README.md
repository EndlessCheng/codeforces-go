[视频讲解](https://www.bilibili.com/video/BV1KG4y1j73o/) 第二题。

用一个最大堆模拟。循环 $k$ 次，每次循环把堆顶加入答案，然后把堆顶 `top` 更新为 `ceil(top/3)`。

原地堆化（heapify）可以做到 $\mathcal{O}(1)$ 的空间复杂度。

部分语言用的标准库自带的堆化函数，具体实现可以看下面的 Java 代码。

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

```java [sol-Java]
class Solution {
    public long maxKelements(int[] nums, int k) {
        heapify(nums); // 原地堆化（最大堆）
        long ans = 0;
        while (k-- > 0) {
            ans += nums[0]; // 堆顶
            nums[0] = (nums[0] + 2) / 3;
            sink(nums, 0); // 堆化（只需要把 nums[0] 下沉）
        }
        return ans;
    }

    // 原地堆化（最大堆）
    // 堆化可以保证 h[0] 是堆顶元素，且 h[i] >= max(h[2*i+1], h[2*i+2])
    private void heapify(int[] h) {
        // 下标 >= h.length / 2 的元素是二叉树的叶子，无需下沉
        // 倒着遍历，从而保证 i 的左右子树一定是堆，那么 sink(h, i) 就可以把左右子树合并成一个堆
        for (int i = h.length / 2 - 1; i >= 0; i--) {
            sink(h, i);
        }
    }

    // 把 h[i] 不断下沉，直到 i 的左右儿子都 <= h[i]
    private void sink(int[] h, int i) {
        int n = h.length;
        while (2 * i + 1 < n) {
            int j = 2 * i + 1; // i 的左儿子
            if (j + 1 < n && h[j + 1] > h[j]) { // i 的右儿子比 i 的左儿子大
                j++;
            }
            if (h[j] <= h[i]) { // 说明 i 的左右儿子都 <= h[i]，停止下沉
                break;
            }
            swap(h, i, j); // 下沉
            i = j;
        }
    }

    // 交换 h[i] 和 h[j]
    private void swap(int[] h, int i, int j) {
        int tmp = h[i];
        h[i] = h[j];
        h[j] = tmp;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxKelements(vector<int> &nums, int k) {
        priority_queue<int> pq(less<int>(), move(nums));
        long long ans = 0;
        while (k--) {
            int x = pq.top();
            pq.pop();
            ans += x;
            pq.push((x + 2) / 3);
        }
        return ans;
    }
};
```

```cpp [sol-C++ 写法二]
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
		ans += int64(h.IntSlice[0]) // 堆顶
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

```js [sol-JavaScript]
var maxKelements = function (nums, k) {
    heapify(nums); // 堆化
    let ans = 0;
    while (k--) {
        ans += nums[0]; // 堆顶
        nums[0] = Math.floor((nums[0] + 2) / 3);
        sink(nums, 0); // 堆化（只需要把 nums[0] 下沉）
    }
    return ans;
};

// 原地堆化（最大堆）
// 堆化可以保证 h[0] 是堆顶元素，且 h[i] >= max(h[2*i+1], h[2*i+2])
function heapify(h) {
    // 下标 >= h.length / 2 的元素是二叉树的叶子，无需下沉
    // 倒着遍历，从而保证 i 的左右子树一定是堆，那么 sink(h, i) 就可以把左右子树合并成一个堆
    for (let i = Math.floor(h.length / 2) - 1; i >= 0; i--) {
        sink(h, i);
    }
}

// 把 h[i] 不断下沉，直到 i 的左右儿子都 <= h[i]
function sink(h, i) {
    const n = h.length;
    while (2 * i + 1 < n) {
        let j = 2 * i + 1; // i 的左儿子
        if (j + 1 < n && h[j + 1] > h[j]) { // i 的右儿子比 i 的左儿子大
            j++;
        }
        if (h[j] <= h[i]) { // 说明 i 的左右儿子都 <= h[i]，停止下沉
            break;
        }
        [h[i], h[j]] = [h[j], h[i]]; // 下沉
        i = j;
    }
}
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
