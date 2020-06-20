package ArrayListDemoDemo

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	array *ArrayListDemo
	capsize int
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.capsize = 10
	stack.array = NewArrayListDemo()
	return stack
}

func (stack *Stack)Clear(){
	stack.capsize = 10
	stack.array.Clear()
}
func (stack *Stack)Size() int{
	size := stack.array.theSize
	return size
}
func (stack *Stack)Pop() interface{}{
	if ! stack.IsEmpty() {
		return nil
	} else {
		value, _ := stack.array.Get(stack.array.theSize - 1)
		stack.array.Delete(stack.array.theSize - 1)
		stack.array.theSize--
		return value
	}
}
func (stack *Stack)Push(data interface{}){
	if !stack.IsFull() {
		stack.array.Append(data)
		stack.array.theSize++
	}
}
func (stack *Stack)IsFull() bool{
	if stack.array.theSize == stack.capsize {
		return true
	} else {
		return false
	}
}
func (stack *Stack)IsEmpty() bool{
	if stack.array.theSize == 0 {
		return true
	} else {
		return false
	}
}
