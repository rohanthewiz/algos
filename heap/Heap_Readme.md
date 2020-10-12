```go
// Heap - Full binary tree - all the levels are full except the lowest node
// The rule of heaps is that the parent must have  a greater value than its children

// Parent - Index * 2 + 1 = left child idx
// Parent - Index * 2 + 2 = right child idx

// Given max heap array
// {50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7}

// Let's consider the node and its children given as value(index):
// parent, left, right
// 48(2), 34(5), 20(6)

// Child index from parent index
p * 2 + 1 = l // example: 2 * 2 + 1 --> 5 is left index
p * 2 + 2 = r // example: 2 * 2 + 2 --> 6 is right index

// Parent index from Left child index
p * 2 + 1 = l
p * 2 = l - 1
p = (l - 1) / 2
// Example: (5 - 1) / 2 --> 2 is parent index

// Parent index from right child index
p * 2 + 2 = r
p * 2 = r - 2
p = (r - 2) / 2
// Example: (6 - 2) / 2 --> 2 is parent index

// Parent index from any child index
// We know that for the left child we get to the parent exactly by subtracting 1 then dividing by 2
// For the right child we can get to the parent exactly by subtracting 2 then dividing by 2 (right is one ahead of left)
// However if we subtracted only one in the case of right we would still get to the same parent as the remainder will be dropped in integer division
// So, we can use a single function (what we have for left) to get to the parent from any child
p = (c - 1) / 2
// Example: (6 - 1) / 2 = 5 / 2 = 2 (remainder discarded)

// HEAP EXTRACT

// - Take the max from index 0
// - Move the last item in the array to the top (index 0)
// - Correct the heap structure from top down
```