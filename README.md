# Data Structures on Go

In this repository I plan to implemment data structres in Go. The code is based on the book
by [Kommadi - Learn data structures and algorithms with golang](https://www.packtpub.com/product/learn-data-structures-and-algorithms-with-golang/9781789618501).


## Linked List

Linked lists is a linear and dynamic data structure that each element points to another element, for instance:

1 --> 2 --> 5 --> 83 --> 190 --> (null)

Assuming that we have N elements in the linked list, we have the following time complexity for operations:

    - Find the last Node: O(N), or O(1) if we have a reference to the end of the list.
    - Add to end: O(N), or O(1) if we have a reference to the end of the list.
    - Find a node with a given property: O(N).
    - Add a node after a given property: O(N).

## Binary Search Tree

A binary search tree is a non-linear data structure that sorts elements that are greater than a given
value on the left branch and less than a given value on the right branch. The space complexity of this
structure is O(N) and time complexity of searching in this structure is O(log N), where N is the number
of elements on the binary tree.

## Queues

Queues are a set of elements ordered by priority, thus elements with higher priority comes first.