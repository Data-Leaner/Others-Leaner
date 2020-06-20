package LinkStack
type LinkNode struct {
	data interface{}
	pNext *LinkNode
}

func NewLinkStack() *LinkNode {
	return &LinkNode{
		data:  nil,
		pNext: nil,
	}
}

type ILinkNode interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Size() int
}

func (linkNode *LinkNode)IsEmpty() bool {
	if linkNode.pNext == nil {
		return true
	} else {
		return false
	}
}
func (linkNode *LinkNode)Push(data interface{}){
	// 头插
	node := &LinkNode{
		data:data,
		pNext: nil,
	}
	node.pNext = linkNode.pNext
	linkNode.pNext = node
}
func (linkNode *LinkNode)Pop() interface{}{
	if linkNode.IsEmpty() {
		return nil
	} else {
		data := linkNode.pNext.data
		linkNode.pNext = linkNode.pNext.pNext
		return data
	}
}
func (linkNode *LinkNode)Size() int{
	var length int
	for linkNode.pNext != nil {
		length++
		linkNode = linkNode.pNext
	}
	return length
}

