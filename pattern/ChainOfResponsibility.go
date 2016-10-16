package pattern

import "fmt"

// 责任链模式
//      使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。
//      将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止。
//
//    github.com/BPing/Golib/middleware 包就是基于此模式实现的


type ScreenEvent struct {
	Type string
	Comment string
}

type IScreenEventHandler interface{
	Handle(*ScreenEvent)bool
	SetNextHandler(IScreenEventHandler)
}

type AbsScreenEventHandler struct {
	NextHandler IScreenEventHandler
}

func(ase *AbsScreenEventHandler)Handle(se *ScreenEvent)bool{
	if ase.NextHandler!=nil {
		return ase.NextHandler.Handle(se)
	}
	return false
}

func(ase *AbsScreenEventHandler)SetNextHandler(ise IScreenEventHandler){
	ase.NextHandler=ise
}

type HomeScreenEventHandler  struct {
	AbsScreenEventHandler
}

func(hse *HomeScreenEventHandler)Handle(se *ScreenEvent)bool{
	fmt.Println("HomeScreenEventHandler.....")
	if se.Type=="HomeClick" {
		fmt.Println("HomeClick")
		return true
	}
	return hse.AbsScreenEventHandler.Handle(se)
}

type UserScreenEventHandler  struct {
	AbsScreenEventHandler
}

func(use *UserScreenEventHandler)Handle(se *ScreenEvent)bool{
	fmt.Println("UserScreenEventHandler.....")
	if se.Type=="UserModelClick" {
		fmt.Println("UserModelClick")
		return true
	}
	return use.AbsScreenEventHandler.Handle(se)
}

func ChainOfResponsibilityTest(){
	  var osd IScreenEventHandler
	   osd=&AbsScreenEventHandler{}
	   home:=&HomeScreenEventHandler{}
	   user:=&UserScreenEventHandler{}
	   home.SetNextHandler(user)
	   osd.SetNextHandler(home)

	screenEvent:=&ScreenEvent{Type:"HomeClick"}
	osd.Handle(screenEvent)

	fmt.Println("-----------------------------------------------\n")
	screenEvent=&ScreenEvent{Type:"UserModelClick"}
	osd.Handle(screenEvent)

	fmt.Println("-----------------------------------------------\n")
	screenEvent=&ScreenEvent{Type:"Null"}
	osd.Handle(screenEvent)
}