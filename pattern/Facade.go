package pattern

import "fmt"

//
// 外观模式
//       隐藏系统的复杂性，并向客户端提供了一个客户端可以访问系统的接口
//

type Coder struct {

}

func (c *Coder)Coding(){
	fmt.Println("Coding ....")
}

type Tester struct {

}

func (t *Tester)Testing(){
	fmt.Println("Testing ....")
}

type ProductPlanner struct {

}

func (p *ProductPlanner)Planing(){
	fmt.Println("Planing ....")
}

type MaintenancePeople struct {

}

func (m *MaintenancePeople)Releasing(){
	fmt.Println("Releasing ....")
}

// 公司
//  拥有产品策划、程序员、测试人员，运维人员
//  通过公司这个外观对外提供服务，
//  而不是直接通过某个类型人员对外服务（虽然最终提供服务的也是某个类型的人员）
type Company  struct {
	ProductPlanner
	Coder
	Tester
	MaintenancePeople
}

// 对外提供生产产品服务
// 需要各个类型人员合作，但细节不对外公开的
func (com *Company)Producing(){
	// 策划产品
	 com.ProductPlanner.Planing()
	// 编码实现
	 com.Coder.Coding()
	// 测试产品
	 com.Tester.Testing()
	// 发布产品
	 com.MaintenancePeople.Releasing()
}