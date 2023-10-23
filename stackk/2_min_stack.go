package stackk

// https://leetcode.com/problems/min-stack/

type MinStack struct {
	size     int
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		0,
		[]int{},
		[]int{},
	}
}

func (minStack *MinStack) Push(val int) {
	minStack.stack = append(minStack.stack, val)
	if minStack.size > 0 && minStack.minStack[minStack.size-1] <= val {
		minStack.minStack = append(minStack.minStack, minStack.minStack[minStack.size-1])
	} else {
		minStack.minStack = append(minStack.minStack, val)
	}
	minStack.size++
}

func (minStack *MinStack) Pop() {
	if minStack.size == 0 {
		return
	}
	minStack.stack = minStack.stack[:minStack.size-1]
	minStack.minStack = minStack.minStack[:minStack.size-1]
	minStack.size--
}

func (minStack *MinStack) Top() int {
	return minStack.stack[minStack.size-1]
}

func (minStack *MinStack) GetMin() int {
	return minStack.minStack[minStack.size-1]
}
