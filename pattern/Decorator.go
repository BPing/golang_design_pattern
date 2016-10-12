package pattern

import (
	"strings"
	"fmt"
)

//
// 装饰模式
//    动态地给一个对象添加一些额外的职责，同时又不改变其结构
//

// 接口
type  MessageBuilder interface{
	 Build(messages ... string) string
}

// 基本信息构造器
type BaseMessageBuilder struct {

}

func(b *BaseMessageBuilder)Build(messages ... string) string{
	return strings.Join(messages,",")
}

// 引号装饰器
type QuoteMessageBuilderDecorator struct {
	Builder MessageBuilder
}

func(q *QuoteMessageBuilderDecorator)Build(messages ... string) string{
	return "\""+q.Builder.Build(messages...)+"\""
}


// 大括号装饰器
type BraceMessageBuilderDecorator struct {
	Builder MessageBuilder
}

func(b *BraceMessageBuilderDecorator)Build(messages ... string) string{
	return "{"+b.Builder.Build(messages...)+"}"
}


// 调试
func DecoratorTest(){
         var MB MessageBuilder

	MB=&BaseMessageBuilder{}

	fmt.Println(MB.Build("hello world"))

	MB=&QuoteMessageBuilderDecorator{MB}

	fmt.Println(MB.Build("hello world"))

	MB=&BraceMessageBuilderDecorator{MB}

	fmt.Println(MB.Build("hello world"))
}