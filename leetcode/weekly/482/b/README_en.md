### Case 1: Each Buys Separately

Without using $\textit{costBoth}$, buy type 1 and type 2 items separately. The total cost is:

$$
\textit{cost}_1 \cdot \textit{need}_1 + \textit{cost}_2 \cdot \textit{need}_2
$$

### Case 2: Use Only Type 3 Items

Buy $\max(\textit{need}_1,\textit{need}_2)$ type 3 items so that both demands are satisfied. The total cost is:

$$
\textit{costBoth} \cdot \max(\textit{need}_1, \textit{need}_2)
$$

### Case 3: Mixed Strategy

If $\textit{costBoth}$ is less than $\textit{cost}_1 + \textit{cost}_2$ but greater than one of them, you can first buy $\min(\textit{need}_1,\textit{need}_2)$ type 3 items, then purchase type 1 or type 2 items to cover the remaining demand.

Assume $\textit{need}_1 \le \textit{need}_2$ (otherwise swap), the total cost is:

$$
\textit{costBoth} \cdot \textit{need}_1 + \textit{cost}_2 \cdot (\textit{need}_2 - \textit{need}_1)
$$

The minimum among these three cases gives the smallest total cost to satisfy all demands.

```py [sol-Python3]
class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        res1 = cost1 * need1 + cost2 * need2  # Each Buys Separately
        if need1 > need2:
            need1, need2 = need2, need1
            cost2 = cost1
        res2 = costBoth * need2  # Use Only Type 3 Items
        res3 = costBoth * need1 + cost2 * (need2 - need1)  # Mixed strategy
        return min(res1, res2, res3)
```

```java [sol-Java]
class Solution {
    public long minimumCost(int cost1, int cost2, int costBoth, int need1, int need2) {
        long res1 = (long) cost1 * need1 + (long) cost2 * need2; // Each Buys Separately
        if (need1 > need2) {
            int tmp = need1;
            need1 = need2;
            need2 = tmp;
            cost2 = cost1;
        }
        long res2 = (long) costBoth * need2; // Use Only Type 3 Items
        long res3 = (long) costBoth * need1 + (long) cost2 * (need2 - need1); // Mixed strategy
        return Math.min(res1, Math.min(res2, res3));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(int cost1, int cost2, int costBoth, int need1, int need2) {
        long long res1 = 1LL * cost1 * need1 + 1LL * cost2 * need2; // Each Buys Separately
        if (need1 > need2) {
            swap(need1, need2);
            cost2 = cost1;
        }
        long long res2 = 1LL * costBoth * need2; // Use Only Type 3 Items
        long long res3 = 1LL * costBoth * need1 + 1LL * cost2 * (need2 - need1); // Mixed strategy
        return min({res1, res2, res3});
    }
};
```

```go [sol-Go]
func minimumCost(cost1, cost2, costBoth, need1, need2 int) int64 {
	res1 := cost1*need1 + cost2*need2 // Each Buys Separately
	if need1 > need2 {
		need1, need2 = need2, need1
		cost2 = cost1
	}
	res2 := costBoth * need2 // Use Only Type 3 Items
	res3 := costBoth*need1 + cost2*(need2-need1) // Mixed strategy
	return int64(min(res1, res2, res3))
}
```

#### Complexity Analysis

- **Time Complexity:** $\mathcal{O}(1)$.
- **Space Complexity:** $\mathcal{O}(1)$.
