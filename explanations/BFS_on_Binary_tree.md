# Explanations by AI using my code as example

## BFS on a Binary Tree

> Source: `algo/BST-binaryTree.go`

### The misconception: "moving between levels without following pointers"

BFS **does** follow pointers. It follows the exact same `node.Left` / `node.Right`
pointers as any other traversal. The queue isn't a replacement for pointers — it's a
**waiting line** that controls the _order_ in which you follow them.

### The core idea: the queue is a "to-do list"

When you pop a node, you don't immediately recurse into its children. Instead you
**drop its children at the back of the queue** and move on:

```go
if node.Left != nil {
    queue.Push(node.Left)
}
if node.Right != nil {
    queue.Push(node.Right)
}
```

Because a queue is FIFO (first-in, first-out), everything you add stays _behind_
everything already waiting. That ordering is what naturally produces level-by-level
output.

The queue refills itself as you go. Each node you take _out_ puts its 0, 1, or 2
children _in_. The queue always holds the **frontier**: nodes discovered but not yet
processed. A node is pushed by its parent, waits its turn, and when it finally gets
popped it becomes the parent that pushes the next batch.

### Why levels stay separated — the snapshot trick

This is the part that makes it work cleanly:

```go
breadth := queue.Size()   // how many nodes are on THIS level right now
for range breadth {        // process exactly that many, no more
```

At the top of each outer loop you take a **snapshot** of how many nodes are in the
queue at that instant. Those are exactly the nodes of the current level. The inner
loop processes _only that many_ — even though you're pushing the next level's children
onto the same queue while you loop, those new children sit at the back and the counter
stops before reaching them. They wait until the next outer iteration, when `depth` has
been bumped.

That's why you can confidently print `depth`: every node popped within one inner loop
belongs to the same level.

### Trace

```
        1
       / \
      2   3
     / \   \
    4   5   6
```

| Step    | `breadth` snapshot | Pop & print     | Push children   | Queue after |
| ------- | ------------------ | --------------- | --------------- | ----------- |
| start   | —                  | —               | push 1          | `[1]`       |
| depth 0 | **1**              | print 1         | push 2, 3       | `[2,3]`     |
| depth 1 | **2**              | print 2, then 3 | push 4,5 then 6 | `[4,5,6]`   |
| depth 2 | **3**              | print 4,5,6     | none            | `[]`        |

### Summary

- **Pointers** = how you _discover_ nodes (parent → children).
- **Queue (FIFO)** = how you _delay_ visiting them, so siblings/cousins get visited
  before any grandchildren.
- **The size snapshot** = how you draw the line between one level and the next.

The loop stops when the queue drains completely (`queue.Size() > 0` becomes false) —
every discovered node has been processed and no new children were pushed.

If you used a _stack_ (LIFO) instead of a queue, you'd dive straight down to a child
before its siblings — that's depth-first. Same pointers, different waiting-line
discipline, completely different order.
