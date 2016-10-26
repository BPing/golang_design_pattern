package pattern

//
// 迭代器模式（Iterator Pattern）
//  这种模式用于顺序访问集合对象的元素，不需要知道集合对象的底层表示。

// 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// 集合接口
type Container interface {
	GetIterator() Iterator
}

// 数组迭代器
type ArrayIterator struct {
	currentIndex int
	ac           ArrayContainer
}

func (ai *ArrayIterator) HasNext() bool {
	if ai.ac.arrayData!=nil&&ai.currentIndex < len(ai.ac.arrayData) {
		return true
	}
	return false
}

func (ai *ArrayIterator) Next() interface{} {
	if ai.HasNext() {
		defer func() { ai.currentIndex++ }()
		return ai.ac.arrayData[ai.currentIndex]

	}
	return nil
}

// 数组集合
type ArrayContainer struct {
	arrayData []interface{}
}

func (ac *ArrayContainer) GetIterator() Iterator {
	return &ArrayIterator{currentIndex: 0, ac: ac}
}


func IteratorTest(){
	arr:=[]string{"a","b","c","d"}
	arrayContainer:=&ArrayContainer{arrayData:arr}
	for iterater:=arrayContainer.GetIterator().(ArrayIterator);iterater.Has {
		iterator
	}
}