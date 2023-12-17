[视频讲解](https://www.bilibili.com/video/BV1994y1A7oo/)

注意本题的子数组不要求是连续的。

既然元素的顺序并不重要，我们应当尽量把相近的数字都放在一起。

所以把数组排序后，从小到大三个三个地切分即可。

> 注：题目保证数组长度是 $3$ 的倍数。

```py [sol-Python3]
class Solution:
    def divideArray(self, nums: List[int], k: int) -> List[List[int]]:
        nums.sort()
        ans = []
        for i in range(2, len(nums), 3):
            if nums[i] - nums[i - 2] > k:
                return []
            ans.append(nums[i - 2: i + 1])
        return ans
```

```java [sol-Java]
public class Solution {
    public int[][] divideArray(int[] nums, int k) {
        Arrays.sort(nums);
        int n = nums.length;
        int[][] ans = new int[n / 3][3];
        for (int i = 2; i < n; i += 3) {
            if (nums[i] - nums[i - 2] > k) {
                return new int[][]{};
            }
            ans[i / 3] = new int[]{nums[i - 2], nums[i - 1], nums[i]};
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> divideArray(vector<int> &nums, int k) {
        sort(nums.begin(), nums.end());
        vector<vector<int>> ans;
        for (int i = 2; i < nums.size(); i += 3) {
            if (nums[i] - nums[i - 2] > k) {
                return {};
            }
            ans.push_back({nums[i - 2], nums[i - 1], nums[i]});
        }
        return ans;
    }
};
```

```go [sol-Go]
func divideArray(nums []int, k int) (ans [][]int) {
	slices.Sort(nums)
	for i := 2; i < len(nums); i += 3 {
		if nums[i]-nums[i-2] > k {
			return nil
		}
		ans = append(ans, nums[i-2:i+1])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
