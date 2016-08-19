package pattern

import "fmt"

//
// 组合模式 将对象组合成树形结构以表示“部分-整体”的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性。
//

type Department struct {
	List   map[string]*Department
	Parent *Department
	Name   string
	Id     string
}

// 批量添加子节点
func (this *Department) Add(departments ...*Department) {
	for _, department := range departments {
		department.AddTo(this)
	}
}

// 添加到对应的父节点
func (this *Department) AddTo(parent *Department) {
	this.Parent = parent
	parent.add(this)
}

// 添加子节点
func (this *Department) add(department *Department) {
	this.List[department.Id] = department
}

// 查找子节点
func (this *Department) Find(id string) (department *Department) {
	department, ok := this.List[id]
	if !ok {
		for _, de := range this.List {
			department = de.Find(id)
			if nil != department {
				return
			}
		}
		return
	}
	return
}

// 移除子节点
func (this *Department) Remove(id string) (department *Department) {
	department, ok := this.List[id]
	if !ok {
		for _, de := range this.List {
			department = de.Remove(id)
			if nil != department {
				return
			}
		}
		return
	}
	delete(this.List, id)
	return
}

// 遍历
func (this *Department) ReadList() {
	fmt.Printf("Node :%s(%s) \n the  children: \n", this.Id, this.Name)
	for _, de := range this.List {
		de.ReadList()
	}
}

// example
func CompositeTest() {
	root := &Department{Name: "root", Id: "0010"}
	children1 := &Department{Name: "network", Id: "00010"}
	children2 := &Department{Name: "forward", Id: "00011"}
	root.Add(children1, children2)
	children1.Add(&Department{Name: "network1", Id: "0001010"})

	root.Find("0001010").ReadList()

	root.ReadList()

	root.Remove("00010").ReadList()

	root.ReadList()

}
