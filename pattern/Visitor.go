package pattern

import (
	"fmt"
	"math"
)

// 访问者模式（Visitor Pattern） 主要将数据结构与数据操作分离
//
// 1、对象结构中对象对应的类很少改变，但经常需要在此对象结构上定义新的操作。
// 2、需要对一个对象结构中的对象进行很多不同的并且不相关的操作，而需要避免让这些操作"污染"这些对象的类，也不希望在增加新操作时修改这些类。


type Visitor interface {
	Visit(DataStruct)
}

type DataStruct interface {
	Accept(Visitor)
}


type ABData struct {
	A int
	B int
}
func (as *ABData)Accept(vi Visitor){
	vi.Visit(as)
}

type AddVisitor struct {

}

func (av *AddVisitor)Visit(dataS DataStruct){
	data:=dataS.(*ABData)
	sum:=data.A+data.B
	fmt.Println("A+B=",sum)
}

type SubVisitor struct {

}

func (sv *SubVisitor)Visit(dataS DataStruct){
	data:=dataS.(*ABData)
	sum:=data.A-data.B
	fmt.Println("abs(A-B)=",math.Abs(float64(sum)))
}

func VisitorTest(){
	Data:=&ABData{A:8,B:10}
	add:=&AddVisitor{}
	sub:=&SubVisitor{}

	Data.Accept(add)
	Data.Accept(sub)
}