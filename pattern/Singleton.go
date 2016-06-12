package pattern

import "fmt"

type Singleton interface {
	SaySomething()
}

type singleton struct {

}

func (singleton)SaySomething() {
	fmt.Println("Singleton")
}

var singletonInstance Singleton

func NewSingletonInstance() Singleton {
	if (nil == singletonInstance) {
		singletonInstance = &singleton{}
	}
	return singletonInstance
}