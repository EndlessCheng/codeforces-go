本题 [视频讲解](https://www.bilibili.com/video/BV14a411U7QZ) 已出炉，额外讲解了**应对设计题的一些技巧**。欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

### 方法一：平衡树

由于数据范围很大，我们可以用一个**哈希表** $\textit{fs}$ 记录每个食物名称对应的食物评分和烹饪方式，另一个**哈希表套平衡树** $\textit{cs}$ 记录每个烹饪方式对应的食物评分和食物名字集合。

对于 `changeRating` 操作，先从 $\textit{cs}[\textit{fs}[\textit{food}].\textit{cuisine}]$ 中删掉旧数据，然后将 $\textit{newRating}$ 和 $\textit{food}$ 记录到 $\textit{cs}$ 和 $\textit{fs}$ 中。

```py [sol1-Python3]
from sortedcontainers import SortedSet

class FoodRatings:
    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.fs = {}
        self.cs = defaultdict(SortedSet)
        for f, c, r in zip(foods, cuisines, ratings):
            self.fs[f] = [r, c]
            self.cs[c].add((-r, f))

    def changeRating(self, food: str, newRating: int) -> None:
        r, c = self.fs[food]
        s = self.cs[c]
        s.remove((-r, food))  # 移除旧数据
        s.add((-newRating, food))  # 添加新数据
        self.fs[food][0] = newRating

    def highestRated(self, cuisine: str) -> str:
        return self.cs[cuisine][0][1]
```

```java [sol1-Java]
class FoodRatings {
    Map<String, Pair<Integer, String>> fs = new HashMap<>();
    Map<String, TreeSet<Pair<Integer, String>>> cs = new HashMap<>();

    public FoodRatings(String[] foods, String[] cuisines, int[] ratings) {
        for (int i = 0; i < foods.length; i++) {
            String f = foods[i], c = cuisines[i];
            var r = ratings[i];
            fs.put(f, new Pair<>(r, c));
            cs.computeIfAbsent(c, k -> new TreeSet<>((a, b) -> !Objects.equals(a.getKey(), b.getKey()) ? b.getKey() - a.getKey() : a.getValue().compareTo(b.getValue()))).add(new Pair<>(r, f));
        }
    }

    public void changeRating(String food, int newRating) {
        var e = fs.get(food);
        var s = cs.get(e.getValue());
        s.remove(new Pair<>(e.getKey(), food)); // 移除旧数据
        s.add(new Pair<>(newRating, food)); // 添加新数据
        fs.put(food, new Pair<>(newRating, e.getValue()));
    }

    public String highestRated(String cuisine) {
        return cs.get(cuisine).first().getValue();
    }
}
```

```cpp [sol1-C++]
class FoodRatings {
    unordered_map<string, pair<int, string>> fs;
    unordered_map<string, set<pair<int, string>>> cs;
public:
    FoodRatings(vector<string> &foods, vector<string> &cuisines, vector<int> &ratings) {
        for (int i = 0; i < foods.size(); ++i) {
            auto &f = foods[i], &c = cuisines[i];
            int r = ratings[i];
            fs[f] = {r, c};
            cs[c].emplace(-r, f);
        }
    }

    void changeRating(string food, int newRating) {
        auto &[r, c] = fs[food];
        auto &s = cs[c];
        s.erase({-r, food}); // 移除旧数据
        s.emplace(-newRating, food); // 添加新数据
        r = newRating;
    }

    string highestRated(string cuisine) {
        return cs[cuisine].begin()->second;
    }
};
```

```go [sol1-Go]
type pair struct {
	r int
	s string
}

type FoodRatings struct {
	fs map[string]pair
	cs map[string]*redblacktree.Tree
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	fs := map[string]pair{}
	ct := map[string]*redblacktree.Tree{}
	for i, f := range foods {
		r, c := ratings[i], cuisines[i]
		fs[f] = pair{r, c}
		if ct[c] == nil {
			ct[c] = redblacktree.NewWith(func(x, y interface{}) int {
				a, b := x.(pair), y.(pair)
				if a.r != b.r {
					return utils.IntComparator(b.r, a.r)
				}
				return utils.StringComparator(a.s, b.s)
			})
		}
		ct[c].Put(pair{r, f}, nil)
	}
	return FoodRatings{fs, ct}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.fs[food]
	t := r.cs[p.s]
	t.Remove(pair{p.r, food}) // 移除旧数据
	t.Put(pair{newRating, food}, nil) // 添加新数据
	p.r = newRating
	r.fs[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	return r.cs[cuisine].Left().Key.(pair).s
}
```

## 方法二：堆

另一种做法是用堆：

- 对于 `changeRating` 操作，直接往 $\textit{cs}$ 中记录，不做任何删除操作；
- 对于 `highestRated` 操作，查看堆顶的食物评分是否等于其实际值，若不相同则意味着对应的元素已被替换成了其他值，堆顶存的是个垃圾数据，直接弹出堆顶；否则堆顶就是答案。

```py [sol2-Python3]
class FoodRatings:
    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.fs = {}
        self.cs = defaultdict(list)
        for f, c, r in zip(foods, cuisines, ratings):
            self.fs[f] = [r, c]
            heappush(self.cs[c], (-r, f))

    def changeRating(self, food: str, newRating: int) -> None:
        f = self.fs[food]
        heappush(self.cs[f[1]], (-newRating, food))  # 直接添加新数据，后面 highestRated 再删除旧的
        f[0] = newRating

    def highestRated(self, cuisine: str) -> str:
        h = self.cs[cuisine]
        while -h[0][0] != self.fs[h[0][1]][0]:  # 堆顶的食物评分不等于其实际值
            heappop(h)
        return h[0][1]
```

```java [sol2-Java]
class FoodRatings {
    Map<String, Pair<Integer, String>> fs = new HashMap<>();
    Map<String, Queue<Pair<Integer, String>>> cs = new HashMap<>();

    public FoodRatings(String[] foods, String[] cuisines, int[] ratings) {
        for (int i = 0; i < foods.length; i++) {
            String f = foods[i], c = cuisines[i];
            var r = ratings[i];
            fs.put(f, new Pair<>(r, c));
            cs.computeIfAbsent(c, k -> new PriorityQueue<>((a, b) -> !Objects.equals(a.getKey(), b.getKey()) ? b.getKey() - a.getKey() : a.getValue().compareTo(b.getValue()))).add(new Pair<>(r, f));
        }
    }

    public void changeRating(String food, int newRating) {
        var c = fs.get(food).getValue();
        cs.get(c).offer(new Pair<>(newRating, food)); // 直接添加新数据，后面 highestRated 再删除旧的
        fs.put(food, new Pair<>(newRating, c));
    }

    public String highestRated(String cuisine) {
        var q = cs.get(cuisine);
        while (!Objects.equals(q.peek().getKey(), fs.get(q.peek().getValue()).getKey())) // 堆顶的食物评分不等于其实际值
            q.poll();
        return q.peek().getValue();
    }
}
```

```cpp [sol2-C++]
class FoodRatings {
    unordered_map<string, pair<int, string>> fs;
    unordered_map<string, priority_queue<pair<int, string>, vector<pair<int, string>>, greater<>>> cs;
public:
    FoodRatings(vector<string> &foods, vector<string> &cuisines, vector<int> &ratings) {
        for (int i = 0; i < foods.size(); ++i) {
            auto &f = foods[i], &c = cuisines[i];
            int r = ratings[i];
            fs[f] = {r, c};
            cs[c].emplace(-r, f);
        }
    }

    void changeRating(string food, int newRating) {
        auto &[r, c] = fs[food];
        cs[c].emplace(-newRating, food); // 直接添加新数据，后面 highestRated 再删除旧的
        r = newRating;
    }

    string highestRated(string cuisine) {
        auto &q = cs[cuisine];
        while (-q.top().first != fs[q.top().second].first) // 堆顶的食物评分不等于其实际值
            q.pop();
        return q.top().second;
    }
};
```

```go [sol2-Go]
type pair struct {
	r int
	s string
}

type FoodRatings struct {
	fs map[string]pair
	cs map[string]*hp
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	fs := map[string]pair{}
	cs := map[string]*hp{}
	for i, f := range foods {
		r, c := ratings[i], cuisines[i]
		fs[f] = pair{r, c}
		if cs[c] == nil {
			cs[c] = &hp{}
		}
		heap.Push(cs[c], pair{r, f})
	}
	return FoodRatings{fs, cs}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.fs[food]
	heap.Push(r.cs[p.s], pair{newRating, food}) // 直接添加新数据，后面 highestRated 再删除旧的
	p.r = newRating
	r.fs[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	h := *r.cs[cuisine]
	for h.Len() > 0 && h[0].r != r.fs[h[0].s].r { // 堆顶的食物评分不等于其实际值
		heap.Pop(&h)
	}
	return h[0].s
}

type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.r > b.r || a.r == b.r && a.s < b.s }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```
