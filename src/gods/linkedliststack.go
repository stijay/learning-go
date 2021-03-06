package examples

import lls "github.com/emirpasic/gods/stacks/linkedliststack"

// LinkedListStackExample to demonstrate basic usage of LinkedListStack
func LinkedListStackExample() {
  stack := lls.New()  // empty
  stack.Push(1)       // 1
  stack.Push(2)       // 1, 2
  stack.Values()      // 2, 1 (LIFO order)
  _, _ = stack.Peek() // 2,true
  _, _ = stack.Pop()  // 2, true
  _, _ = stack.Pop()  // 1, true
  _, _ = stack.Pop()  // nil, false (nothing to pop)
  stack.Push(1)       // 1
  stack.Clear()       // empty
  stack.Empty()       // true
  stack.Size()        // 0