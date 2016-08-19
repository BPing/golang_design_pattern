package pattern

//
//另一种单例模式的实现方式，也是golang语言常用的模式
//采用init（）方法全局之实例化一次
//

import "fmt"

type Singleton2 interface {
	SaySomething()
}

type singleton2 struct {
}

func (singleton2) SaySomething() {
	fmt.Println("Singleton")
}

var singletonInstance2 Singleton2

func init() {
	singletonInstance2 = new(singleton2)
}

func Singleton2SaySomething() {
	singletonInstance2.SaySomething()
}
