## Core Idea

1. Consider each bit of the answer from high to low, deciding whether it can be $1$ or must be $0$.
2. Traverse $\textit{nums}$ and compute the minimum number of operations needed to turn certain bits of $\textit{nums}[i]$ into $1$.
3. Compute the sum $s$ of the smallest $m$ operation counts. If $s ≤ k$, then this bit of the answer can be $1$; otherwise, it must be $0$.

## Detailed Explanation

Since the result of the AND operation does not exceed $\max(\textit{nums}) + k$, we can start enumerating bits from the binary length of $\max(\textit{nums}) + k$ minus $1$.

Suppose we are currently checking whether the AND result can be $target$, and let $x = \textit{nums}[i]$.

For example, if the binary number $target = 0100$, this means that for every $\textit{nums}[i]$, the third bit (counting from low to high) must be turned into $1$.

For example:

- If $x = 0001$, then $x$ needs to be increased to $0100$.
- If $x = 1010$, then $x$ needs to be increased to $1100$.

In general, we find the highest bit (from high to low) where $target$ is $1$ and $x$ is $0$. The bits of $x$ higher than this position remain unchanged, and the remaining lower bits are increased so that they match $target$.

Using this method, we compute the required number of operations for each $\textit{nums}[i]$, then take the sum $s$ of the smallest $m$ operation counts. If $s ≤ k$, then this bit of the answer can be $1$; otherwise, it must be $0$.

⚠ **Note**: For example, if $target = 0100$ and this bit of the answer is determined to be $1$, then in the next iteration, the $target$ to be checked is $0110$, not $0010$. **Bits of the answer that have already been determined to be $1$ cannot be changed.**

```py [sol-Python3]
class Solution:
    def maximumAND(self, nums: List[int], k: int, m: int) -> int:
        ops = [0] * len(nums)  # Number of operations for each number
        ans = 0
        max_width = (max(nums) + k).bit_length()
        for bit in range(max_width - 1, -1, -1):
            target = ans | (1 << bit)  # Note: target includes the bits already set in ans
            for i, x in enumerate(nums):
                j = (target & ~x).bit_length()
                # j-1 is the highest bit where target is 1 and x is 0
                mask = (1 << j) - 1
                ops[i] = (target & mask) - (x & mask)

            # Greedy: pick the smallest m operation counts
            ops.sort()
            if sum(ops[:m]) <= k:
                ans = target  # This bit of the answer can be set to 1
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumAND(int[] nums, int k, int m) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] ops = new int[nums.length]; // Number of operations for each number
        int ans = 0;
        int maxWidth = 32 - Integer.numberOfLeadingZeros(mx + k);
        for (int bit = maxWidth - 1; bit >= 0; bit--) {
            int target = ans | (1 << bit); // Note: target includes the bits already set in ans
            for (int i = 0; i < nums.length; i++) {
                int x = nums[i];
                int j = 32 - Integer.numberOfLeadingZeros(target & ~x);
                // j-1 is the highest bit where target is 1 and x is 0
                int mask = (1 << j) - 1;
                ops[i] = (target & mask) - (x & mask);
            }

            // Greedy: pick the smallest m operation counts
            Arrays.sort(ops);
            long sum = 0;
            for (int i = 0; i < m; i++) {
                sum += ops[i];
            }
            if (sum <= k) {
                ans = target; // This bit of the answer can be set to 1
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumAND(vector<int>& nums, int k, int m) {
        vector<int> ops(nums.size()); // Number of operations for each number
        int ans = 0;
        int max_width = bit_width((uint32_t) ranges::max(nums) + k);
        for (int bit = max_width - 1; bit >= 0; bit--) {
            int target = ans | (1 << bit); // Note: target includes the bits already set in ans
            for (int i = 0; i < nums.size(); i++) {
                int x = nums[i];
                int j = bit_width((uint32_t) target & ~x);
                // j-1 is the highest bit where target is 1 and x is 0
                int mask = j < 31 ? (1 << j) - 1 : INT_MAX;
                ops[i] = (target & mask) - (x & mask);
            }

            // Greedy: pick the smallest m operation counts
            // ranges::sort(ops);
            ranges::nth_element(ops, ops.begin() + m);
            if (reduce(ops.begin(), ops.begin() + m, 0LL) <= k) {
                ans = target; // This bit of the answer can be set to 1
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumAND(nums []int, k, m int) (ans int) {
	ops := make([]int, len(nums)) // Number of operations for each number
	maxWidth := bits.Len(uint(slices.Max(nums) + k))
	for bit := maxWidth - 1; bit >= 0; bit-- {
		target := ans | 1<<bit // Note: target includes the bits already set in ans
		for i, x := range nums {
			j := bits.Len(uint(target &^ x))
			// j-1 is the highest bit where target is 1 and x is 0
			mask := 1<<j - 1
			ops[i] = target&mask - x&mask
		}

		// Greedy: pick the smallest m operation counts
		slices.Sort(ops)
		sum := 0
		for _, x := range ops[:m] {
			sum += x
		}
		if sum <= k {
			ans = target // This bit of the answer can be set to 1
		}
	}
	return
}
```

#### Complexity Analysis

- Time Complexity: $\mathcal{O}(n\log n\log U)$ or $\mathcal{O}(n\log U)$, where $n$ is the length of $\textit{nums}$ and $U = \max(\textit{nums}) + k$. Using a quick-select algorithm, it can be reduced to $\mathcal{O}(n\log U)$, as shown in the C++ code.
- Space Complexity: $\mathcal{O}(n)$.
