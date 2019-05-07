package main

import (
	"fmt"
	"reflect"
)

var a int
var b int=456

func main() {
	a=123
	fmt.Println(a)
	fmt.Println(b)
	var (
		c string="慕课网"
		d int=789
	)
	fmt.Println(c,d)
	//var x,y,z  int=1,2,3
	//var x,y,z=1,2.1,3
	//q,w,e :=11,12,13     //只能在函数体内，不能用在全局变量上
	//fmt.Println(q,w,e)
	//fmt.Println(reflect.TypeOf(x))
	//fmt.Println(reflect.TypeOf(y))
	//fmt.Println(x,y,z)
	//var m,_,n =21,22,23   //下划线
	//fmt.Println(m,n)
	//var x int =3
	var y float32  =3.01
	//fmt.Println(x)
	//z :=float32(x)
	z :=int32(y)
	//fmt.Println(z,reflect.TypeOf(z))
	fmt.Println(z,reflect.TypeOf(z))

	//常量
	const name string ="simon"     //显示
	fmt.Println(name)
	const name1 ="xuming"          //隐示
	fmt.Println(name1)
}
