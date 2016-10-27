package pattern

import "fmt"

// 状态模式（State Pattern）
//   核心思想就是：当对象的状态改变时，同时改变其行为，很好理解！
// 就拿QQ来说，有几种状态，在线、隐身、忙碌等，每个状态对应不同的操作，
// 而且你的好友也能看到你的状态， 所以，状态模式就两点：
//   1、可以通过改变状态来获得不同的行为。
//   2、你的好友能同时看到你的变化


// 相关模式：命令模式

// 灯
type Light struct {
     State LightState
}

func(l *Light)PressSwitch(){
	if l.State!=nil {
		l.State.PressSwitch(l)
	}
}

// 灯状态
type LightState  interface {
	PressSwitch(light *Light)
}


type OnLightState struct {

}

// 开灯
func(ols *OnLightState)PressSwitch(light *Light){
	fmt.Println("turn on light")
	// 下一个状态动作为关灯
	light.State=&OffLightState{}
}

type OffLightState struct {

}

// 关灯
func(ols *OffLightState)PressSwitch(light *Light){
	fmt.Println("turn off light")
	// 下一个状态动作为开灯
	light.State=&OnLightState{}
}

func StateTest(){
	light:=&Light{State:&OnLightState{}}
	light.PressSwitch()
	light.PressSwitch()
	light.PressSwitch()
	light.PressSwitch()
}