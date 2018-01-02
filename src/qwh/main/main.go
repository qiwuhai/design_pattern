package main

import (
	"fmt"
	"qwh/common"
)

func main() {
	oper := common.Calculate{Operation: "+"}
	result := oper.Cal(1, 3)
	fmt.Println(result)
}
