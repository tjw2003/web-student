# GO语言基础

### 1.简单程序

```go
package main

import "fmt"

func main() {
    fmt.Printf("Hello, world or 你好，世界\n")
}
```

输出：

```
Hello, world or 你好，世界
```

Go 程序是通过 `package` 来组织的,`package <pkgName>`（在我们的例子中是 `package main` ）这一行告诉我们当前文件属于哪个包，而包名 `main` 则告诉我们它是一个可独立运行的包，它在编译后会产生可执行文件。除了 `main` 包之外，其它的包最后都会生成 `*.a` 文件。

我们调用了一个函数 `Printf`，这个函数来自于 `fmt` 包，所以我们在第三行中导入了系统级别的 `fmt` 包：`import "fmt"`。

在第五行中，我们通过关键字 `func` 定义了一个 `main` 函数，函数体被放在 `{}`（大括号）中.

第六行，我们调用了 `fmt` 包里面定义的函数 `Printf`。大家可以看到，这个函数是通过 `<pkgName>.<funcName>` 的方式调用的.

### 定义变量

使用 `var` 关键字是 Go 最基本的定义变量方式，与 C 语言不同的是 Go 把变量类型放在变量名后面

```go
// 定义一个名称为 “variableName”，类型为 "type" 的变量
var variableName type
// 定义三个类型都是 “type” 的变量
var vname1, vname2, vname3 type
// 初始化 “variableName” 的变量为 “value” 值，类型是 “type”
var variableName type = value
/*
    定义三个类型都是 "type" 的变量, 并且分别初始化为相应的值
    vname1 为 v1，vname2 为 v2，vname3 为 v3
*/
var vname1, vname2, vname3 type= v1, v2, v3
/*
    定义三个变量，它们分别初始化为相应的值
    vname1 为 v1，vname2 为 v2，vname3 为 v3
    编译器会根据初始化的值自动推导出相应的类型
*/
vname1, vname2, vname3 := v1, v2, v3
//数组
var arr [n]type
arr[0] = 42      // 数组下标是从0开始的
c := [...]int{4, 5, 6} // 可以省略长度而采用 `...` 的方式，Go 会自动根据元素个数来计算长度
```

在 Go 中，布尔值的类型为 `bool`，值是 `true` 或 `false`，默认为 `false`。

常量const

slice与map

`slice` 并不是真正意义上的动态数组，而是一个引用类型。`slice` 总是指向一个底层 `array`，`slice` 的声明也可以像 `array` 一样，只是不需要长度。

```go
// slice和声明 array 一样，只是少了长度
var fslice []int
slice := []byte {'a', 'b', 'c', 'd'}//slice初始化示例

//slice 可以从一个数组或一个已经存在的 slice 中再次声明。slice 通过 array[i:j] 来获取，其中 i 是数组的开始位置，j 是结束位置，但不包含 array[j]，它的长度是 j-i。

// 声明一个含有 10 个元素元素类型为 byte 的数组
var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

// 声明两个含有 byte 的 slice
var a, b []byte

// a 指向数组的第 3 个元素开始，并到第五个元素结束，
a = ar[2:5]
// 现在 a 含有的元素: ar[2]、ar[3]和ar[4]

// b 是数组 ar 的另一个 slice
b = ar[3:5]
// b 的元素是：ar[3] 和 ar[4]
 
```

`map` 的读取和设置也类似 `slice` 一样，通过 `key` 来操作，只是 `slice` 的 `index` 只能是 `int` 类型，而 `map` 多了很多类型

```go
// 声明一个 key 是字符串，值为 int 的字典, 这种方式的声明需要在使用之前使用 make 初始化
var numbers map[string]int
// 另一种 map 的声明方式
numbers := make(map[string]int)
numbers["one"] = 1  // 赋值
numbers["ten"] = 10 // 赋值
numbers["three"] = 3

fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
// 打印出来如:第三个数字是: 3
```



### 流程和函数

##### if

Go 里面 `if` 条件判断语句中不需要括号

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is less than 10")
}

// 计算获取值 x, 然后根据 x 返回的大小，判断是否大于 10。
if x := computedValue(); x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is less than 10")
}

```

##### for

```go
package main

import "fmt"

func main(){
    sum := 0;
    for index:=0; index < 10 ; index++ {
        sum += index
    }
    fmt.Println("sum is equal to ", sum)
}
// 输出：sum is equal to 45

```

##### switch

```go
switch sExpr {
case expr1:
    some instructions
case expr2:
    some other instructions
case expr3:
    some other instructions
default:
    other code
}

```

##### 函数

```go
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    // 这里是处理逻辑代码
    // 返回多个值
    return value1, value2
}



//示例
package main

import "fmt"

// 返回 a、b 中最大值.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    x := 3
    y := 4
    z := 5

    max_xy := max(x, y) // 调用函数max(x, y)
    max_xz := max(x, z) // 调用函数max(x, z)

    fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
    fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
    fmt.Printf("max(%d, %d) = %d\n", y, z, max(y,z)) // 也可在这直接调用它
}

```

##### 多个返回值

```go
package main

import "fmt"

// 返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
    return A+B, A*B
}

func main() {
    x := 3
    y := 4

    xPLUSy, xTIMESy := SumAndProduct(x, y)

    fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
    fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}

```

##### 变参

```go
func myfunc(arg ...int) {}
//arg ...int 告诉 Go 这个函数接受不定数量的参数。注意，这些参数的类型全部是 int。在函数体中，变量 arg 是一个 int 的 slice
```

##### 指针

```go
package main

import "fmt"

// 简单的一个函数，实现了参数+1的操作
func add1(a *int) int { // 请注意，
    *a = *a+1 // 修改了a的值
    return *a // 返回新值
}

func main() {
    x := 3

    fmt.Println("x = ", x)  // 应该输出 "x = 3"

    x1 := add1(&x)  // 调用 add1(&x) 传x的地址

    fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
    fmt.Println("x = ", x)    // 应该输出 "x = 4"
}

```

