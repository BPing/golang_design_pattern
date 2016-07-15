package pattern

//
// 组合模式 将对象组合成树形结构以表示“部分-整体”的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性。
//

type Department struct {
	List []*Department
	Name string
	Id   string
}

func (d *Department)Add(departments... *Department) {
	d.List = append(d.List, departments ...)
}
func (d *Department)Remove(id string) {


}
func (d *Department)List() {}

