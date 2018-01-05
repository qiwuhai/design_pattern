package main

import (
	"fmt"
	"qwh/common"
)

func main() {
	var ope string
	var param1, param2 float64
	fmt.Println("enter two number:")
	fmt.Scanf("%f %f", &param1, &param2)
	fmt.Println("enter operation:")
	fmt.Scanln(&ope)

	factory := common.NewFactoryCal()
	oper := factory.CreateOper(ope)
	oper.SetData(param1, param2)
	result := oper.Operater()
	fmt.Println(result)
}
