package pattern

import "fmt"

// 模板模式（Template Pattern）
//   定义一个操作中算法的框架，而将一些步骤延迟到子类中。
// 模板方法模式使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

type Game interface {
	Start()
	Playing()
	End()
}

type AbsGame struct {

}

func(ag *AbsGame)Start(){
	fmt.Println("start game")
}
func(ag *AbsGame)Playing(){
	fmt.Println("game playing")
}
func(ag *AbsGame)End(){
	fmt.Println("the end")
}

type FootBall struct {
	*AbsGame
}

func(fb *FootBall)Playing(){
	fmt.Println("playing FootBall")
}

type Cricket  struct {
	*AbsGame
}

func(fb *Cricket )Playing(){
	fmt.Println("playing Cricket ")
}

func RunGame(g Game){
	g.Start()
	g.Playing()
	g.End()
}

func TemplateTest(){
	RunGame(&FootBall{})
	RunGame(&Cricket{})
}