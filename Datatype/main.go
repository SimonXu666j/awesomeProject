package main

import (
	"awesomeProject/test"
	"fmt"
	"unsafe"
)
//别名
type 慕课网 int32
func main() {
	//无符号整型
	//var i uint8 =1
	//var a int32  =1
	//var b int  = 1
	//var c int  = 1
	//var d float32  =1.0
	//var e float64  =1.0
	//var bl bool=true
	//var f byte  = 1
	//var g rune  =1
	//fmt.Println(unsafe.Sizeof(i))
	//fmt.Println(unsafe.Sizeof(a))
	//fmt.Println(unsafe.Sizeof(b))
	//fmt.Println(unsafe.Sizeof(c))
	//fmt.Println(unsafe.Sizeof(d))
	//fmt.Println(unsafe.Sizeof(e))
	//fmt.Println(bl)
	//fmt.Println(unsafe.Sizeof(f))
	//fmt.Println(unsafe.Sizeof(g))
	//各类数值类型默认值0
	//var i int32
	//var j float32
	//var b bool
	//var d complex64
	//fmt.Println(i,j,b,d)
	//字符串类型默认值为空值
	//var s string
	//fmt.Println(s)
	//var bm 慕课网=1
	//fmt.Println(bm,reflect.TypeOf(bm),unsafe.Sizeof(bm))
	bt :=byte(1)
	fmt.Println(unsafe.Sizeof(bt))
	i :=int32(123)
	fmt.Println(unsafe.Sizeof(i))
	j :=float32(5.16)
	fmt.Println(unsafe.Sizeof(j))
	a :=int64(123)
	fmt.Println(unsafe.Sizeof(a))
	name :=[5]byte{'i','m','o','o','c'}
	fmt.Println(unsafe.Sizeof(name))
	num :=[4]int{121,223,323}
	fmt.Println(unsafe.Sizeof(num))
	fmt.Println(test.Car)
	//fmt.Println(test.a)
}
