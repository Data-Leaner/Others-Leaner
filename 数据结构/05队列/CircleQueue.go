package CircleQueue

const QueueSize = 100

type ICircleQueue interface {
	Size() int
	Front() interface{}
	End() interface{}
	IsEmpty() bool
	IsFull() bool
	EnQueue(data interface{})
	DeQueue() interface{}
	Clear()
}

type CircleQueue struct {
	dataStore []interface{}
	front int
	rear int
}

func NewCircleQueue() *CircleQueue {
	circleQueue := new(CircleQueue)
	circleQueue.dataStore = make([]interface{},QueueSize,QueueSize)
	circleQueue.front = 0
	circleQueue.rear = 0
	return circleQueue
}

func (circleQueue *CircleQueue)Size() int {
	return (circleQueue.rear - circleQueue.front + QueueSize) % QueueSize
}
func (circleQueue *CircleQueue)Front() interface{}{
	if circleQueue.IsEmpty() {
		return nil
	} else {
		return circleQueue.dataStore[circleQueue.front]
	}
}
func (circleQueue *CircleQueue)End() interface{}{
	if circleQueue.IsEmpty() {
		return nil
	} else {
		return circleQueue.dataStore[circleQueue.rear]
	}
}
func (circleQueue *CircleQueue)IsEmpty() bool{
	if circleQueue.front == circleQueue.rear {
		return true
	} else {
		return false
	}
}
func (circleQueue *CircleQueue)IsFull() bool{
	if (circleQueue.rear + 1)%QueueSize == (circleQueue.front)%QueueSize {
		return true
	} else {
		return false
	}
}
func (circleQueue *CircleQueue)EnQueue(data interface{}){
	if circleQueue.IsFull() {
		return
	} else {
		circleQueue.dataStore[circleQueue.rear] = data
		circleQueue.rear = (circleQueue.rear + 1) % QueueSize
	}

}
func (circleQueue *CircleQueue)DeQueue() interface{}{
	if circleQueue.IsEmpty() {
		return nil
	} else {
		frontValue := circleQueue.dataStore[circleQueue.front]
		circleQueue.dataStore[circleQueue.front] = 0
		circleQueue.front = (circleQueue.front + 1) % QueueSize
		return frontValue
	}
}
func (circleQueue *CircleQueue)Clear(){
	circleQueue.dataStore = make([]interface{},QueueSize,QueueSize)
	circleQueue.front = 0
	circleQueue.rear = 0
}
