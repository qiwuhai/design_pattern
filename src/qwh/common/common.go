package common

import (
	"fmt"
)

type Calculate struct {
	Operation string `json:"operation"`
}

func (this *Calculate) Cal(a int, b int) int {
	oper := this.Operation
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("unknow operation")
	}
	return 0
}
