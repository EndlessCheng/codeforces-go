For ease of calculation, change the initial sequence to start from $0$, i.e., $0,1,2,\ldots,n-1$.

In the first operation, we delete all odd numbers, leaving only even numbers. This means that, for the final answer (starting from $0$), the least significant bit in binary must be $0$.

Right-shift all remaining elements $0,2,4,\ldots$ by one bit, and we obtain the sequence $0,1,2,\ldots$ again. Based on this, perform the second operation.

In the second operation, we delete from right to left:

- If the last number of the sequence is even, for example $0,1,2,3,4$, then we will delete all odd numbers, leaving only even numbers.
- If the last number of the sequence is odd, for example $0,1,2,3,4,5$, then we will delete all even numbers, leaving only odd numbers.

This means that, for the final answer (starting from $0$), the second bit from low to high in binary must be equal to the second bit from low to high of $n-1$.

And so on.

In general, for the final answer (starting from $0$), the binary bits at positions $0,2,4,\ldots$ from low to high must be $0$, while the bits at positions $1,3,5,\ldots$ must be the same as those of $n-1$ at the corresponding positions.

After computing the answer, add one to it (because the original sequence starts from $1$).

```py [sol-Python3]
class Solution:
    def lastInteger(self, n: int) -> int:
        MASK = 0xAAAAAAAAAAAAAAA  # ...1010
        return ((n - 1) & MASK) + 1
```

```java [sol-Java]
class Solution {
    public long lastInteger(long n) {
        final long MASK = 0xAAAAAAAAAAAAAAAL; // ...1010
        return ((n - 1) & MASK) + 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long lastInteger(long long n) {
        constexpr long long MASK = 0xAAAAAAAAAAAAAAALL; // ...1010
        return ((n - 1) & MASK) + 1;
    }
};
```

```go [sol-Go]
func lastInteger(n int64) int64 {
	const mask = 0xAAAAAAAAAAAAAAA // ...1010
	return (n-1)&mask + 1
}
```

#### Complexity Analysis

**Time complexity**: $\mathcal{O}(1)$.

**Space complexity**: $\mathcal{O}(1)$.
