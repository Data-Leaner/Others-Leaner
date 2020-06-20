package Queue

type IQueue interface {
	Size() int
	Front() interface{}
	End() interface{}
	IsEmpty() bool
	IsFull() bool
	EnQueue(data interface{})
	DeQueue() interface{}
	Clear()
}

type Queue struct {
	// 队列的数据存储
	dataStore []interface{}
	// 队列的大小
	theSize int
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.dataStore = make([]interface{}, 0, 10)
	queue.theSize = 0
	return queue
}

func (queue *Queue)Size() int{
	return queue.theSize
}
func (queue *Queue)Front() interface{}{
	if queue.IsEmpty() {
		return nil
	} else {
		frontValue := queue.dataStore[0]
		return frontValue
	}
}
func (queue *Queue)End() interface{}{
	if queue.IsEmpty() {
		return nil
	} else {
		endValue := queue.dataStore[queue.theSize-1]
		return endValue
	}
}
func (queue *Queue)IsEmpty() bool{
	if queue.theSize == 0 {
		return true
	} else {
		return false
	}
}
func (queue *Queue)IsFull() bool{
	if queue.theSize == cap(queue.dataStore) {
		return true
	} else {
		return false
	}
}

func (queue *Queue)EnQueue(data interface{}){
	if queue.IsFull() {
		return
	} else {
		queue.dataStore = append(queue.dataStore, data)
		queue.theSize++
	}
}
func (queue *Queue)DeQueue() interface{}{
	if queue.IsEmpty() {
		return nil
	} else {
		deQueueValue := queue.dataStore[0]
		queue.dataStore = append(queue.dataStore[1:queue.theSize])
		queue.theSize--
		return deQueueValue
	}
}
func (queue *Queue)Clear(){
	queue.dataStore = make([]interface{}, 0, 10)
	queue.theSize = 0
}


