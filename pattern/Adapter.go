package pattern

import "fmt"

//
// 适配器模式
//    是作为两个不兼容的接口之间的桥梁。这种类型的设计模式属于结构型模式，它结合了两个独立接口的功能。
//    或者
//    将一个类的接口转换成客户希望的另外一个接口。
//    Adapter 模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
//
//   主要分为三类：类的适配器模式、对象的适配器模式、接口的适配器模式
//


//不可充电电池使用接口
type NonRechargeableBattery interface {
	Use()
}

//可充电电池使用接口
type RechargeableBattery interface {
	Use()
	Charge()
}

//不可充电电池A
type  NonRechargeableA struct {
}

func (NonRechargeableA)Use() {
	fmt.Println("NonRechargeableA using ")
}

//类的适配器模式
// 似乎golang语言层面已经支持了，暂且搁置
//
//


//对象的适配器模式

//适配可充电电池使用接口
type AdapterNonToYes struct {
	NonRechargeableBattery
}

func (AdapterNonToYes)Charge() {
	fmt.Println("AdapterNonToYes Charging")
	//nothing to do ,just adapt for RechargeableBattery's interface
}


//接口的适配器模式

type RechargeableBatteryAbstract struct {

}

func (RechargeableBatteryAbstract) Use() {
	fmt.Println("RechargeableBatteryAbstract using")
}
func (RechargeableBatteryAbstract) Charge() {
	fmt.Println("RechargeableBatteryAbstract Charging")
}

type NonRechargeableB struct {
	RechargeableBatteryAbstract
}

func (NonRechargeableB) Use() {
	fmt.Println("NonRechargeableB using ")
}


//test
func AdapterTest() {
	var battery RechargeableBattery

	battery = AdapterNonToYes{NonRechargeableA{}}
	battery.Use()
	battery.Charge()

	battery = NonRechargeableB{}
	battery.Use()
	battery.Charge()
}