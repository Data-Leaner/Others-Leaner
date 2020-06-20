package StackArray

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	dataSource []interface{}
	// 最大范围
	capsize int
	// 当前数据
	currentSize int
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.dataSource = make([]interface{},0, 10)
	stack.capsize = 10
	stack.currentSize = 0
	return stack
}

func (stack *Stack)Clear(){
	stack.dataSource = make([]interface{},0, 10)
	stack.capsize = 10
	stack.currentSize = 0
}
func (stack *Stack)Size() int{
	return stack.currentSize
}
func (stack *Stack)Pop() interface{}{
	if stack.IsEmpty() {
		return nil
	} else {
		lastValue := stack.dataSource[stack.currentSize-1]
		stack.dataSource = stack.dataSource[:stack.currentSize-1]
		stack.currentSize--
		return lastValue
	}
}
func (stack *Stack)Push(data interface{}){
	if stack.IsFull() {

	} else {
		stack.dataSource = append(stack.dataSource, data)
		stack.currentSize++
	}

}
func (stack *Stack)IsFull() bool{
	if stack.currentSize >= stack.capsize {
		return true
	} else {
		return false
	}

}
func (stack *Stack)IsEmpty() bool{
	if stack.currentSize == 0 {
		return true
	} else {
		return false
	}
}
