package pattern

//抽象工厂
type AbstractFactory interface {
	Produce() (pen) //生产笔
}

type PencilFactory struct {

}

func (PencilFactory)Produce() (pen) {
	return new(pencil)
}

type BrushPen struct {

}

func (BrushPen)Produce() (pen) {
	return new(brushPen)
}