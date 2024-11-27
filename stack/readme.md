
# Introduction

The stack is a fundamental data structure that functions using the principle of `last in` - `first out`, often
abbreviated LIFO.

You can imagine the stack as a stack of books, for flags:

```shell
[🇺🇦]
[🇷🇴]
[🇺🇸]
```
The stack's Push (🇩🇪) operation adds an element to the stack and it will look like this:

```shell
[🇩🇪]
[🇺🇦]
[🇷🇴]
[🇺🇸]
```

The stack can also implement the following operations:

- Pop: Removes an element from the stack
- Peek: Peeks at the top element of the stack.
- Size: Returns the stack size.

Since the stack is often implemented using an Array data structure, for convenience
the LIFO principle doesn't always need to hold. Implementers can also provide operations
for iterating through the stack, converting it to a list, checking if an element is present into the stack
retrieving the first element and so on.

## Use Cases

The stack has several use cases:

### Tracking functions and storing data
Interpreters use a stack to track the function calls of programs. 

When a function calls and returns nested function's value the
nested function is added to the stack, in case it also has a nested function in it, which is also added to the stack.
The interpreter executes the top function to get its value and uses it to execute function from the rest of the stack.

### Undo

A stack can be used to implement an undo mechanism.

### Graph Algorithms

Stacks are used to keep the state when executing an algorithm on a graph.

### Memory Allocation

In many programming languages, local variables are stored on the stack.

# Implementation

To implement a stack with an array, you will need two fields:

- The array which stores the data.
- A field to keep track of the stack size.

```go
// Stack is a simple stack implementation that holds uint16 data.
type Stack[T any] struct {
	storage   []T
	stackSize uint32
}
```

## New

The new method creates a new instance of the stack:

```go
// NewStack creates a new Stack instance.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		storage:   make([]T, 0, 128),
		stackSize: 0,
	}
}
```

## Push

The push operation pushes the data onto the stack. At the top.

```go

// Push pushes data into the stack.
func (s *Stack[T]) Push(data T) {
	s.stackSize += 1
	s.storage = append(s.storage, data)
}
```

## Pop

The pop operation returns the data from the stack and removes it.

```go
// Pop pops an element from the stack.
func (s *Stack[T]) Pop() (T, error) {
	var whatever T
	if s.stackSize == 0 {
		return whatever, errors.New("stack is empty")
	}
	data := s.storage[s.stackSize-1]
	s.stackSize -= 1
	return data, nil
}
```

## Peek

The peek operation returns the data from the stack without removing it.

```go
// Peek peeks at the latest data stored on the stack.
func (s *Stack[T]) Peek() T {
	return s.storage[s.stackSize-1]
}
```

# Analysis

The analysis of the stack implemented with an array is as following.

## Space

The space occupied by the stack is equal to the elements present in the stack.

This means that the stack is O(N). The Go array has an initial capacity and grows as needed.

## Time

- Pop / Peek / Push - Are O(1), because we access the array using the first or last index.
- Size - Is O(1)
- Contains - Is O(N), because we need to iterate through all the elements from the stack to check if it contains the given element.

# Conclusion

The stack is a fundamental and simple data structure. Every programmer benefits by learning it due to its wast use cases.

