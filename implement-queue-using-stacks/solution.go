package implementqueueusingstacks

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) reorder() {
	if len(this.out) > 0 {
		return
	}

	for len(this.in) > 0 {
		i := len(this.in) - 1
		value := this.in[i]
		this.in = this.in[:i]
		this.out = append(this.out, value)
	}
}

func (this *MyQueue) Pop() int {
	this.reorder()

	ret := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return ret
}

func (this *MyQueue) Peek() int {
	this.reorder()

	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
