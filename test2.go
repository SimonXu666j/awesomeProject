//单行注释
//程序所属包
package main
//导入依赖包
import "fmt"
//常量定义
const NAME  string= "Simon"
//全局变量的声明与赋值
var mainName string = "global name"
//一般类型声明
type imoocInt int
//结构体声明
type Learn struct {

}
//声明接口
type Ilearn interface {

}
//函数定义
func learnImooc()  {
	fmt.Print("learnImooc")
}
/*
多行注释
*/
//main函数
func main() {
	fmt.Println("learn1")
	fmt.Println(mainName)
	fmt.Println(NAME)
	learnImooc()
}
