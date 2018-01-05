/*1.定义一个接口Operater，包含两个方法;
    四个运算类型(Add,Sub,Mul,Div)实现该接口
  2.定义FactoryCal工厂类结构体，绑定方法根据输入参数返回具体的运算结构体类型
  3.实现了各个运算模块的独立，方便修改某个运算或添加新的运算类型
  4.调用端只需通过FactoryCal获取所需的运算，然后调用Operater接口方法即可;修改或新增运算类型不影响客户端
  --------------
  客户端调用demo:
  func main() {
	var ope string = "+"
	var param1, param2 float64 = 1.0,2.0
	factory := common.NewFactoryCal()
	operoper.SetData(param1, param2)
	result := oper.Operater()
	fmt.Println(result)
  }
*/
package common

import ()

type Operater interface {
	SetData(float64, float64)
	Operater() float64
}

type BaseOpe struct {
	ParamsA float64
	ParamsB float64
}
type Add struct {
	BaseOpe
}

func (this *Add) Operater() float64 {
	return this.ParamsA + this.ParamsB
}
func (this *Add) SetData(a float64, b float64) {
	this.ParamsA = a
	this.ParamsB = b
}

type Sub struct {
	BaseOpe
}

func (this *Sub) Operater() float64 {
	return this.ParamsA - this.ParamsB
}
func (this *Sub) SetData(a float64, b float64) {
	this.ParamsA = a
	this.ParamsB = b
}

type Mul struct {
	BaseOpe
}

func (this *Mul) Operater() float64 {
	return this.ParamsA * this.ParamsB
}
func (this *Mul) SetData(a float64, b float64) {
	this.ParamsA = a
	this.ParamsB = b
}

type Div struct {
	BaseOpe
}

func (this *Div) Operater() float64 {
	return this.ParamsA / this.ParamsB
}
func (this *Div) SetData(a float64, b float64) {
	this.ParamsA = a
	this.ParamsB = b
}

type FactoryCal struct {
}

func NewFactoryCal() *FactoryCal {
	ins := new(FactoryCal)
	return ins
}
func (f FactoryCal) CreateOper(oper string) Operater {
	switch oper {
	case "+":
		return &Add{}
	case "-":
		return &Sub{}
	case "*":
		return &Mul{}
	case "/":
		return &Div{}
	default:
		panic("invalid")
		return nil
	}
}
