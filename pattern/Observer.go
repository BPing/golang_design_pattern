package pattern

import "fmt"

//
// 观察者模式（Observer Pattern）
//  定义对象间的一种一对多的依赖关系,
//  当一个对象的状态发生改变时, 所有依赖于它的对象都得到通知并被自动更新。
//

type Observer interface {
	Notify(interface{})
}

type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject)SetState(state string) {
	s.state = state
	s.NotifyAllObservers()
}
func (s *Subject)Attach(observer ... Observer) {
	s.observers = append(s.observers, observer ...)
}

func (s *Subject)NotifyAllObservers() {
	for _, obs := range s.observers {
		obs.Notify(s)
	}
}

type AObserver struct {
	Id string
}

func (ao *AObserver)Notify(sub interface{}) {
	fmt.Println(ao.Id , " receive ", sub.(*Subject).state)
}

func ObserverTest() {
	sub := &Subject{}
	a := &AObserver{Id:"A"}
	b := &AObserver{Id:"b"}
	sub.Attach(a,b)
	sub.SetState("hello world")
	sub.SetState("i know")

}