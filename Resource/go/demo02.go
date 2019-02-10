package main 
import "fmt"

// 变量的使用
func varUse() {
	// 方式一
	var name string 
	name = "CoderHG"

	// 方式二
	var sex = "handsome"

	/** 方式三
	var age
	age = 18
	*/// 这种方式是不对的

	// 方式四 这种方式有点厉害了....
	girlFriend := false

	// 打印结果: CoderHG handsome false
	fmt.Println(name, sex, /**age*/ girlFriend)
}

//  多变量的使用
func varMUse() {
	// 方式一
	var name, sex string = "CoderHG", "handsome"
	// 方式二
	var (
		x int = 10
		y int = 10
		address string = "TJ"
		)

	// 打印结果: CoderHG handsome 10 10 TJ
	fmt.Println(name, sex, x, y, address)
}

// 没有使用已经声明的变量, 编译直接报错
func unUseVar() {
	// var a string = "abc"
}

func constUse() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   // 打印结果: 面积为 : 50
   fmt.Printf("面积为 : %d", area)
   println()
   // 打印结果: 1 false str
   println(a, b, c)   
}

// 函数定义
func funcUse(a int, b int, c string) (int, int, string) {
	return a*10, b*5, c
}

// 数组
func arrayUse() {
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balance[4] = 50.0

	var j = 0
	/* 输出每个数组元素的值 */
   for j = 0; j < 5; j++ {
      fmt.Println("balance[", j, "] = ", balance[j])
   }

   /** 打印结果
		balance[ 0 ] =  1000
		balance[ 1 ] =  2
		balance[ 2 ] =  3.4
		balance[ 3 ] =  7
		balance[ 4 ] =  50
   */
}

// 指针
func pointerUse() {
	var a = 3
	var pointer *int 

	if (pointer != nil) {
		fmt.Println("pointer 有值")
	} else {
		fmt.Println("空指针")
	} // 打印结果: 空指针

	pointer = &a
	if (pointer == nil) {
		fmt.Println("空指针")
	} else {
		fmt.Println("pointer 有值")
	} // 打印结果: pointer 有值

	// 修改 a 的值
	a = 10

	// 打印结果: 10
	fmt.Println(*pointer)
}


// 定义一个结构体
type Books struct {
   title string
   author string
   subject string
   book_id int
}

// 结构体
func structUse() {
	// 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

/** 打印结果:
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com  0}
*/

func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}

// 打印切片
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

// 切片
func sampleSlipeUse() {
	// 创建切片
	// var numbers = make([]int,3,5)
	numbers := []int{0,1,2,3,4,5,6,7,8}
	// 打印切片
	printSlice(numbers)
	// 打印结果: len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

	// 剪切切片
	number1 := numbers[2:]
	number2 := numbers[2:5]
	
	// 打印结果: len=7 cap=7 slice=[2 3 4 5 6 7 8]
	printSlice(number1)
	// 打印结果: len=3 cap=7 slice=[2 3 4]
	printSlice(number2)

	// 同时添加多个元素
	number2 = append(number2, 21,31,41)
	// 打印结果: len=6 cap=7 slice=[2 3 4 21 31 41]
	printSlice(number2)
}

func rangeUse() {
	// 这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    // 在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    // range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    // range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

/** 打印结果
sum: 9
index: 1
a -> apple
b -> banana
0 103
1 111
*/

// map
func mapUse() {
	// 定义一个集合
	var countryCapitalMap map[string]string 
	// 创建集合实例
    countryCapitalMap = make(map[string]string)

	// map插入key - value对,各个国家对应的首都 
    countryCapitalMap [ "France" ] = "Paris"
    countryCapitalMap [ "Italy" ] = "罗马"

	// 使用键输出地图值 
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    /*查看元素在集合中是否存在 */
    captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(captial) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("美国的首都是", captial)
    } else {
        fmt.Println("美国的首都不存在")
    }

    fmt.Println("=====删除元素======")

    // 删除 : delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
    /*删除元素*/ 
    delete(countryCapitalMap, "France")


    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    } // 发现已经没有 France 的数据

}

/** 打印结果:
France 首都是 Paris
Italy 首都是 罗马
美国的首都不存在
=====删除元素======
Italy 首都是 罗马
*/

// 阶乘
func Factorial(n uint64) (uint64) {
	if n > 0 {
		return n * Factorial(n-1);
	}
	return 1
}

// 斐波那契数列
func fibonacci(n int) int {
	if (n < 2) {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

// recursion
func recursionUse() {
	// 阶乘
	var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
    // 打印结果: 15 的阶乘是 1307674368000

	// 斐波那契数列
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i))
    } // 打印结果: 0    1    1    2    3    5    8    13    21    34
}

// 类型转换
func typeChangedUse() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	// 打印结果: mean 的值为: 3.400000
	fmt.Printf("mean 的值为: %f\n",mean)
}


// 定义接口
type Phone interface {
    call()
}

// 结构体定义 诺基亚
type NokiaPhone struct {

}

// 函数实现
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// 结构体定义 苹果
type IPhone struct {

}

// 函数实现
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// 接口
func interfaceUse() {
	var phone Phone

    phone = new(NokiaPhone)
    // 打印结果: I am Nokia, I can call you!
    phone.call()

    phone = new(IPhone)
    // 打印结果: I am iPhone, I can call you!
    phone.call()
}

// 定义一个 DivideError 结构
type DivideError struct {
    dividee int
    divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0`

    return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
        dData := DivideError{
            dividee: varDividee,
            divider: varDivider,
        }
        errorMsg = dData.Error()
        return
    } else {
        return varDividee / varDivider, ""
    }
}

func errorUse() {
	// 正常情况
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
        fmt.Println("100/10 = ", result)
    }
    // 当被除数为零的时候会返回错误信息
    if _, errorMsg := Divide(100, 0); errorMsg != "" {
        fmt.Println("errorMsg is: ", errorMsg)
    }
}

/** 打印结果:
	100/10 =  10
errorMsg is:  
    Cannot proceed, the divider is zero.
    dividee: 100
    divider: 0
*/
func main() {
	// varUse()
	// varMUse()
	// unUseVar()
	// constUse()

	// 函数调用
	// a, b, c := funcUse(10, 8, "99")
	// 打印结果: 100 40 99
	// fmt.Println(a, b, c)

	// 数组
	// arrayUse()
	// pointerUse()

	// structUse()
	// sampleSlipeUse()
	// rangeUse()
	// mapUse()
	// recursionUse()
	// typeChangedUse()
	// interfaceUse()
	errorUse()
}