package main

import (
	"fmt"
	"qwh/common"
)

func main() {
	oper := common.Calculate{Operation: "+"}
	result := oper.Cal(1, 5)
	fmt.Println(result)
}
