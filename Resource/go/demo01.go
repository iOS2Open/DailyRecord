// 引入的包名
package main
// 具体的包
import "fmt"

// 先于 main 函数调用
func init() {
	fmt.Println("INIT")
}

// 主函数
func main() {
	// 简单的变量定义
	var a int
	a = 78
	var b int
	b = 20
	// 定义并做加法运算
	var c int = a + b
	// 简单的打印
	fmt.Println("Hello,  CoderHG!  c = ", c)
}