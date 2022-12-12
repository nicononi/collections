# Go Collections
Golang Collections. This is for learning and practicing purposes and not intended to be used in production.
A working example can be found in the `example` folder

## Stack

Structure that follows the Last-In-First-Out design pattern.

### Usage

```
var stack collections.Stack[int] = new(collections.SliceStack[int])
stack.Push(10)
peek := stack.Peek()
removed := stack.Pop()
```

## Queue

Structure that follows the First-In-First-Out design pattern.

### Usage

```
var queue collections.Queue[int] = new(collections.SliceQueue[int])
queue.Enqueue(10)
peek := stack.Peek()
removed := stack.Dequeue()
```

## List

Structure that applies the Single Linked List pattern

### Usage

```
var list collections.List[int] = new(collections.SliceList[int])
list.Append(10)
peek := stack.Get(0)
contains := stack.Contains(0)
removed := stack.Remove(0)
```