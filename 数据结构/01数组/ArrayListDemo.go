package ArrayListDemoDemo

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	// 数组的大小
	Size() int
	// 抓取第几个元素
	Get(index int) (interface{}, error)
	// 修改元素
	Set(index int, newValue interface{}) error
	// 插入数据
	Insert(index int, newValue interface{}) error
	// 追加函数
	Append(newValue interface{})
	// 清空
	Clear()
	// 删除数据
	Delete(index int) error
	// 返回字符串
	String() string
	//
	Iterator() Iterator
}
// 数据结构
type ArrayListDemo struct {
	// 数组存储
	dataStore []interface{}
	// 数组的大小
	theSize int
}

func NewArrayListDemo() *ArrayListDemo {
	list := new(ArrayListDemo)
	// 开辟空间10个
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (list *ArrayListDemo)checkIsFull() {
	// 判断内存的使用
	if list.theSize == cap(list.dataStore) {
		newDataStore := make([]interface{}, list.theSize * 2, list.theSize * 2)
		copy(newDataStore, list.dataStore)
		list.dataStore = newDataStore
	}
}

func (list *ArrayListDemo)Size() int {
	return list.theSize
}
// 抓取第几个元素
func (list *ArrayListDemo)Get(index int) (interface{}, error){
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	} else {
		return list.dataStore[index], nil
	}
}
// 追加数据
func (list *ArrayListDemo)Append(newValue interface{}){
	list.dataStore = append(list.dataStore, newValue)
	list.theSize++
}
// 修改元素
func (list *ArrayListDemo)Set(index int, newValue interface{}) error{
	if index < 0 || index >= list.theSize {
		return errors.New("数组越界")
	} else {
		list.dataStore[index] = newValue
		return nil
	}
}
// 插入数据
func (list *ArrayListDemo)Insert(index int, newValue interface{}) error{
	// 检测内存，如果空间满了，则自动扩展
	list.checkIsFull()
	list.dataStore = list.dataStore[:list.theSize+1]
	for i := list.theSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newValue
	list.theSize++
	return nil
}
// 清空
func (list *ArrayListDemo)Clear(){
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
}
// 删除数据
func (list *ArrayListDemo)Delete(index int) error{
	if index < 0 || index >= list.theSize {
		return errors.New("数组越界")
	} else {
		list.dataStore = append(list.dataStore[0:index], list.dataStore[index+1:]...)
		list.theSize--
		return nil
	}
}
// 返回字符串
func (list *ArrayListDemo)String() string{
	// 打印成字符串
	return fmt.Sprint(list.dataStore)
}














