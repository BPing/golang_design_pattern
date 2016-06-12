package pattern

//工厂模式适合：凡是出现了大量的产品需要创建，
//并且具有共同的接口时，可以通过工厂方法模式进行创建
import "fmt"

//笔
type pen interface {
	//写字
	Write()
}

type pencil struct {
}

func (p *pencil) Write() {
	fmt.Println("铅笔")
}

type brushPen struct {
}

func (p *brushPen) Write() {
	fmt.Println("毛笔")
}

//工厂
type PenFactory struct {

}

func (this PenFactory)Produce(typ string) (pen) {
	switch typ {
	case "pencil":
		return this.ProducePencil()
	case "brush":
		return this.ProduceBrushPen()
	default:
		return nil
	}
}

func (PenFactory)ProducePencil() (pen) {
	return new(pencil)
}

func (PenFactory)ProduceBrushPen() (pen) {
	return new(brushPen)
}