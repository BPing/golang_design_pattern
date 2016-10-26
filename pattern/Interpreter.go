package pattern

import (
	"strings"
	"fmt"
)

//解释器模式（Interpreter Pattern）
//    提供了评估语言的语法或表达式的方式，文法解释，它属于行为型模式。
// 这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。
//
// 一般递归处理



type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	Word string
}

// 终结符
func(te *TerminalExpression)Interpret(context string) bool{
	if strings.Contains(context,te.Word) {
		return true
	}
	return false
}

// 或
type OrExpression struct {
	A Expression
	B Expression
}

func(oe *OrExpression)Interpret(context string) bool{
	return oe.A.Interpret(context)||oe.B.Interpret(context)
}

// 与
type AndExpression  struct {
	A Expression
	B Expression
}

func(ae *AndExpression )Interpret(context string) bool{
	return ae.A.Interpret(context)&&ae.B.Interpret(context)
}

func InterpreterTest(){
	isMale :=&OrExpression{&TerminalExpression{"Robert"},&TerminalExpression{"John"}}
	isMarriedWoman :=&AndExpression{&TerminalExpression{"Julie"},&TerminalExpression{"Married"}}
	fmt.Println("John is male?",isMale.Interpret("John"))
	fmt.Println("Julie is a married women?",isMarriedWoman.Interpret("Married Julie"))
}