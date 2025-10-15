package implementstackusingqueues

type MyStack struct {
    stack []int
}


func Constructor() MyStack {
    return MyStack{stack: make([]int, 0)}
}


func (this *MyStack) Push(x int)  {
  this.stack = append(this.stack, x)
}


func (this *MyStack) Pop() int {
    l := len(this.stack)
    elem := this.stack[l-1]
    this.stack = this.stack[:l-1]
    return elem
}


func (this *MyStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MyStack) Empty() bool {
    return len(this.stack) == 0
}