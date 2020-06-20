package ArrayListDemoDemo

type Iterator interface {
	HasNext() bool
	Next() interface{}
	Remove()
	GetIndex() int
}

//type Iterable interface {
//	Iterator() Iterator // 构造初始化接口
//}

// 构造指针访问数组
type ArrayListIterator struct {
	list *ArrayListDemo
	// 当前索引
	currentIndex int
}

func (list *ArrayListDemo)Iterator() Iterator{
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator)HasNext() bool {
	if it.currentIndex < it.list.theSize {
		return true
	} else {
		return false
	}
}

func (it *ArrayListIterator) Next() interface{} {
	if !it.HasNext() {
		return nil
	} else {
		getValue, _ := it.list.Get(it.currentIndex)
		it.currentIndex++
		return getValue
	}
}
func (it *ArrayListIterator)Remove(){
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}
func (it *ArrayListIterator)GetIndex() int{

}
