package pattern

const (
	FOOD  = "food"
	DRINK = "drink"
)

type Item interface {
	Price() float32   //价钱
	Name() string     //名称
	Category() string //类别
}

//食物
type Food struct {
}

func (Food) Price() float32 {
	return 0.0
}
func (Food) Name() string {
	return ""
}
func (Food) Category() string {
	return FOOD
}

//饮料
type Drink struct {
}

func (Drink) Price() float32 {
	return 0.0
}
func (Drink) Name() string {
	return ""
}
func (Drink) Category() string {
	return DRINK
}

//汉堡
type Hamburger struct {
	Food
}

func (Hamburger) Price() float32 {
	return 12.00
}
func (Hamburger) Name() string {
	return "Hamburger"
}

//炸鸡
type FriedChicken struct {
	Food
}

func (FriedChicken) Price() float32 {
	return 18.00
}
func (FriedChicken) Name() string {
	return "FriedChicken"
}

//可乐
type Cola struct {
	Drink
}

func (Cola) Price() float32 {
	return 3.00
}
func (Cola) Name() string {
	return "Cola"
}

//啤酒
type Beer struct {
	Drink
}

func (Beer) Price() float32 {
	return 5.00
}
func (Beer) Name() string {
	return "Beer"
}

type Meal []Item

func (this *Meal) AddItem(item ...Item) {
	*this = append(*this, item...)
}

func (this Meal) GetCost() (cost float32) {
	for _, val := range this {
		cost += val.Price()
	}
	return
}
func (this Meal) ShowItems() (msg string) {
	for _, val := range this {
		msg += "Category：" + val.Category() + " Name:" + val.Name() + "\n"
	}
	return
}

//建造者
type MealBuilder struct {
}

func (MealBuilder) MealOne() (meal *Meal) {
	meal = new(Meal)
	meal.AddItem(new(FriedChicken), new(Beer))
	return
}
