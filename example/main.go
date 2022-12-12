package main

import (
	"collections"
	"fmt"
)

func main() {
	testStack()
	testQueue()
	testList()
}

func testStack() {
	fmt.Println("Testing Stack")
	stack := new(collections.SliceStack[int])

	fmt.Println("Pushing 10")
	stack.Push(10)
	fmt.Println("Pushing 20")
	stack.Push(20)
	fmt.Println("Pushing 30")
	stack.Push(30)

	for !stack.IsEmpty() {
		fmt.Println("Poping first element")
		fmt.Println(stack.Pop())
	}
}

func testQueue() {
	fmt.Println("")
	fmt.Println("Testing Queue")
	queue := new(collections.SliceQueue[int])

	fmt.Println("Enqueing 10")
	queue.Enqueue(10)
	fmt.Println("Enqueing 20")
	queue.Enqueue(20)
	fmt.Println("Enqueing 30")
	queue.Enqueue(30)

	for !queue.IsEmpty() {
		fmt.Println("Dequeing first element")
		fmt.Println(queue.Dequeue())
	}
}

func testList() {
	fmt.Println("")
	fmt.Println("Testing List")
	list := new(collections.SliceList[int])

	fmt.Println("Appending 10")
	list.Append(10)
	fmt.Println("Appending 20")
	list.Append(20)
	fmt.Println("Appending 30")
	list.Append(30)

	fmt.Println("Contains 10?")
	fmt.Println(list.Contains(10))
	fmt.Println("Contains 0?")
	fmt.Println(list.Contains(0))

	for !list.IsEmpty() {
		fmt.Println("Removing first element")
		fmt.Println(list.Remove(0))
	}
}
