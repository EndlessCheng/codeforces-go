# codeforces-go 💭💡🎈

## Algorithm

I have implemented some algorithms in [copypasta](./copypasta).

## How to Practice

> 1. Starting from problems with difficulty equal to ("Your rating" + 200), you will be solving problem that will be quite hard for you.
>
> 2. Spend no more than 10 minutes before look at the solution.
>
> 3. Train on constructive problems ("constructive algorithms" tag on CF) will significantly improve the time you spend on finding observations. 
>
> [source](https://codeforces.com/blog/entry/66715?#comment-507869)

## BST

Golang doesn't have a built-in BST. I use [treap](./copypasta/treap.go) after some research.

> Binary search tree (BST) based data structures, such as AVL trees, red-black trees, and splay trees, are often used in system software, such as operating system kernels. 
> Choosing the right kind of tree can impact performance significantly, but the literature offers few empirical studies for guidance. 
> We compare 20 BST variants using three experiments in real-world scenarios with real and artificial workloads. 
> The results indicate that when input is expected to be randomly ordered with occasional runs of sorted order, red-black trees are preferred; 
> when insertions often occur in sorted order, AVL trees excel for later random access, whereas splay trees perform best for later sequential or clustered access. 
> **For node representations, use of parent pointers is shown to be the fastest choice**, with threaded nodes a close second choice that saves memory; nodes without parent pointers or threads suffer when traversal and modification are combined; maintaining an in-order doubly linked list is advantageous when traversal is very common; and right-threaded nodes perform poorly.
>
> See [Performance Analysis of BSTs in System Software](misc/Performance%20Analysis%20of%20BSTs%20in%20System%20Software.pdf) for more detail.

Left: [treap](./copypasta/treap.go) (xorshift32 random number)

Right: [red black tree](./copypasta/red_black_tree.go)

![](misc/bst.png)

## Codeforces Solutions

[main](./main)

## Others

My GoLand `Live Templates` and `Postfix Completion` [settings](./misc/my_goland_template)

Account:

[![](https://cfrating.ihcr.top/?user=synapse)](https://codeforces.com/profile/synapse)

### Useful Tools

[Draw Geometry](https://csacademy.com/app/geometry_widget/)

[Draw Graph](https://csacademy.com/app/graph_editor/)

[OEIS](https://oeis.org/)

[Wolfram|Alpha](https://www.wolframalpha.com/)

[UpSolve.me](https://upsolve.me/)

[Contests Filter](https://codeforceshelper.herokuapp.com/contests)

[Practice Problems Recommender](https://recommender.codedrills.io/)

[Codeforces Upsolving Helper](https://codeforces-upsolving-helper.herokuapp.com/)

[Codeforced](http://codeforced.github.io/handle/)

### Rating

[Open Codeforces Rating System](https://codeforces.com/blog/entry/20762)

[Codeforces Visualizer](https://cfviz.netlify.app/)

[Codeforces: Problem Difficulties](https://codeforces.com/blog/entry/62865)

### Keep Healthy

[Exercises!](https://musclewiki.org/)
