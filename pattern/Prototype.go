package pattern

import (
	"fmt"
	"reflect"
)

//
// 原型模式
//  用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。
//
// ps:或者可以使用sync.Pool来替代原型模式方式
//

type Cloneable interface {
	Clone() interface{}
}

type Context struct {
}

func (this *Context) Clone() interface{} {
	new_obj := (*this)
	return &new_obj
}

//结构体基本信息
func (this *Context) string(typ interface{}) (str string) {
	v := reflect.ValueOf(typ).Elem()
	str += "Type:" + v.Type().Name() + "\n"
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		str += fmt.Sprintf("index %d: %s %s = %v\n", i,
			v.Type().Field(i).Name, f.Type(), f.Interface())
	}
	return
}

type Context1 struct {
	Uri string
	Context
}

func (this *Context1) SetUri(uri string) {
	this.Uri = uri
}

func (this *Context1) String() (str string) {
	return this.string(this)
}

type Context2 struct {
	Context
	Echo string
}

func (this *Context2) SetEcho(echo string) {
	this.Echo = echo
}

func (this *Context2) String() (str string) {
	return this.string(this)
}

//原型池
type ContextPool map[string]*Context

func (this *ContextPool) AddContext(key string, val *Context) {
	(*this)[key] = val
}

func (this *ContextPool) GetContext(key string) *Context {
	val, ok := (*this)[key]
	if ok {
		return val.Clone().(*Context)
	}
	return nil
}
