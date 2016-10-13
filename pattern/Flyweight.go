package pattern

import (
	"fmt"
)

//
// 享元模式
//    实现对象的共享，即共享池，当系统中对象多的时候可以减少内存的开销，通常与工厂模式一起使用。
//    sync.pool 就属于此模式
//

// 数据库连接池
type DbConnectPool struct {
	ConnChan chan *DbConnect
}

func NewDbConnectPool(chanLen int)*DbConnectPool{
	return &DbConnectPool{ConnChan:make(chan *DbConnect,chanLen)}
}

func(dc *DbConnectPool) Get()*DbConnect{
	select {
	case conn := <-dc.ConnChan:
		return conn
	default:
	        // 无则新建
		return new(DbConnect)
	}
}
func(dc *DbConnectPool) Put(conn *DbConnect){
	select {
	case dc.ConnChan <- conn:
		return
	default:
	       // 满则丢弃
		return
	}
}

// 数据库连接
type DbConnect struct {

}

func (*DbConnect) Do(){
	fmt.Println("connect......doing....")
}


func FlyweightTest(){
	pool:=NewDbConnectPool(2)
	conn:=pool.Get()
	conn.Do()
	pool.Put(conn)
}
