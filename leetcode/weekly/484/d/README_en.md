## Core Idea

1. Build the answer bit by bit from high to low, deciding whether each bit is $1$ or $0$.
2. Traverse $\textit{nums}$ and compute the minimum number of operations needed to turn certain bits of $\textit{nums}[i]$ into $1$.
3. Compute the sum $s$ of the smallest $m$ operation counts. If $s ≤ k$, then this bit of the answer can be $1$; otherwise, it must be $0$.

## Detailed Explanation

We construct the answer from the most significant bit to the least significant bit, deciding whether each bit is $1$ or $0$.

> Since the AND result cannot exceed $\max(\textit{nums}) + \lfloor\frac{k}{m}\rfloor$, we start enumerating from the highest bit, i.e. one less than the binary length of $\max(\textit{nums}) + \lfloor\frac{k}{m}\rfloor$.

Suppose we are checking whether the AND result can include $\textit{target}$. For example, let $\textit{target} = 0100$ (in binary). Then, for $x = \textit{nums}[i]$, the third bit from low to high must be $1$ to ensure that the third bit of the AND result is $1$.

For example:

- If $x = 0001$, then $x$ needs to be increased to $0100$.
- If $x = 1010$, then $x$ needs to be increased to $1100$.

In general, we find the **highest missing bit** of $x$: that is, the first bit position $j$ (from high to low) where $\textit{target}$ has $1$ but $x$ has $0$.  

All bits of $x$ higher than $j$ remain unchanged, while the lower $j$ bits are increased so that they become equal to the lower $j$ bits of $\textit{target}$.

Using this method, we compute the required number of operations for each $\textit{nums}[i]$. Then we take the sum $s$ of the smallest $m$ operation counts. If $s ≤ k$, we set this bit of the answer to $1$; otherwise, we set it to $0$.

## Q&A

**Q**: When processing each bit of the answer, the chosen $m$ numbers are not necessarily the same as the $m$ numbers chosen in the previous iteration. What if they differ?

**A**: In each iteration, the algorithm only checks whether it is *possible* to select $m$ numbers that satisfy the requirement. As long as such a selection exists, this bit of the answer can be set to $1$. There is no need to ensure that the selected numbers are identical across iterations. The $m$ numbers chosen in the final iteration where a bit is set to $1$ are the actual final selection. Note that these $m$ numbers satisfy not only the current bit being $1$, but also all higher bits of $target$ that were previously set to $1$.

```py [sol-Python3]
class Solution:
    def maximumAND(self, nums: List[int], k: int, m: int) -> int:
        ops = [0] * len(nums)  # Number of operations for each number
        ans = 0
        max_width = (max(nums) + k // m).bit_length()
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
        int maxWidth = 32 - Integer.numberOfLeadingZeros(mx + k / m);
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
        int max_width = bit_width((uint32_t) ranges::max(nums) + k / m);
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
	maxWidth := bits.Len(uint(slices.Max(nums) + k/m))
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
