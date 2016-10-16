package pattern

import (
	"errors"
	"fmt"
)

// 代理模式
//     为其他对象提供一种代理以控制对这个对象的访问
//     优点： 1、职责清晰。 2、高扩展性。 3、智能化。
//     注意事项：
//          1、和适配器模式的区别：适配器模式主要改变所考虑对象的接口，
//             而代理模式不能改变所代理类的接口。
//          2、和装饰器模式的区别：装饰器模式为了增强功能，而代理模式
//             是为了加以控制。


// 似乎有点像装饰模式

type Device interface {
	Read()([]byte,error)
	Write(word []byte)error
}

type HardDisk struct {
	storage []byte
}

func(h *HardDisk)Read()([]byte,error){
	return h.storage,nil
}
func(h *HardDisk)Write(word []byte)error{
	h.storage=word
	return nil
}

type HardDiskProxy struct {
	OpId string
	hd *HardDisk
}

func(h *HardDiskProxy)Read()([]byte,error){
	if !h.permission("read") {
		return nil ,errors.New("You don't have permission to read")
	}
	return h.hd.Read()
}
func(h *HardDiskProxy)Write(word []byte)error{
	if !h.permission("wrire") {
		return errors.New("You don't have permission to write")
	}
	return h.hd.Write(word)
}

// 权限判断
func(h *HardDiskProxy)permission(tag string)bool{
	if h.OpId=="admin"{
		return true
	}

	if h.OpId=="reader"&&tag=="read"{
		return true
	}

	if h.OpId=="writer"&&tag=="wrire"{
		return true
	}

	return false
}

func ProxyTest(){
	var devI Device
	devI=&HardDiskProxy{OpId:"admin",hd:&HardDisk{}}
	devI.Write([]byte("Hello world!"))
	data,_:=devI.Read()
	fmt.Println(string(data))

	devI=&HardDiskProxy{OpId:"reader",hd:&HardDisk{storage:[]byte("only read")}}
	err:=devI.Write([]byte("Hello world!"))
	fmt.Println(err.Error())
	data,_=devI.Read()
	fmt.Println(string(data))

	devI=&HardDiskProxy{OpId:"writer",hd:&HardDisk{}}
	err=devI.Write([]byte("only writer"))
	fmt.Println(err)
	data,err=devI.Read()
	fmt.Println(string(data),err.Error())
}