package pattern

import (
	"fmt"
)

// 策略模式（Strategy Pattern）
// 定义一系列的算法,把它们一个个封装起来,
// 并且使它们可相互替换。本模式使得算法可独立于使用它的客户而变化。

// 相关模式：工厂模式

type Strategy interface {
	Do(interface{})
}

// 从A到B
type AToB struct {
	// 距离
	ABDistance float64
        // 到达方式策略
	Strategy Strategy
}

func(ab *AToB) Do(){
	if ab.Strategy!=nil {
		ab.Strategy.Do(ab)
	}
}

type BikeStrategy struct {
	Speed float64
}

func(bs *BikeStrategy) Do(ab interface{}){
	aTob,ok:=ab.(*AToB)
	if ok &&bs.Speed<=0.0000001{
		return
	}
	fmt.Println("方式：自行车 用时：",aTob.ABDistance/bs.Speed)
}

type BusStrategy struct {
	Speed float64
}

func(bs *BusStrategy) Do(ab interface{}){
	aTob,ok:=ab.(*AToB)
	if ok &&bs.Speed<=0.0000001{
		return
	}
	fmt.Println("方式：巴士 用时：",aTob.ABDistance/bs.Speed)
}

type AirStrategy struct {
	Speed float64
}

func(as *AirStrategy) Do(ab interface{}){
	aTob,ok:=ab.(*AToB)
	if ok &&as.Speed<=0.0000001{
		return
	}
	fmt.Println("方式：飞机 用时：",aTob.ABDistance/as.Speed)
}

func StrategyTest(){
        aTob:=&AToB{ABDistance:600}

	aTob.Strategy=&BikeStrategy{Speed:15}
	aTob.Do()

	aTob.Strategy=&BusStrategy{Speed:90}
	aTob.Do()

	aTob.Strategy=&AirStrategy{Speed:500}
	aTob.Do()

}