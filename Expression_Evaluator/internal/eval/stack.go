package eval

type (
	node[N any] struct {
		value N
		prev  *node[N]
	}

	stack[T any] struct {
		top    *node[T]
		length int
	}
)

func (this *stack[T]) len() int {
	return this.length
}

func (this *stack[T]) peek() T {
	if this.length == 0 {
		panic(ERROR_EMPTY_STACK_PEEKING)
	}
	return this.top.value
}

func (this *stack[T]) pop() T {
	if this.length == 0 {
		panic(ERROR_EMPTY_STACK_POPPING)
	}

	n := this.top
	this.top = n.prev
	this.length--

	return n.value
}

func (this *stack[T]) push(value T) {
	n := &node[T]{value, this.top}
	this.top = n
	this.length++
}
