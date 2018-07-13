go语言
# 特点，优势

运行效率高，开发效率高，部署简单
语言层面支持并发
内置runtime
内置gofmt

脚本化 
静态类型 + 编译型
原生支持并发编程

# 劣势

语法糖没有python多
运行速度不及c，比c++，java快
第三方库比较少

# 应用

服务器编程
网络编程
分布式系统
云平台


# 环境安装

下载安装：
配置环境：PATH D:\Program Files\go1.9.2\bin
版本检查：>go version
配置：

GOPATH 工作目录
GOROOT golang安装路径


命令行工具
go build 编译源码
go run 编译并运行源码文件
go get 获取远程代码

IDE
GoLand：https://www.jetbrains.com/go/?fromMenu
激活：http://idea.youbbs.org

# 基础语法

关键字
标识符 25个
注释 单行注释//  多行注释 /* */
基础结构
package
import
别名
路径
"."
"_"
变量 函数 可见性规则
数据类型

# package

是最基本的分发单位和工程管理中依赖关系的体现

每个go语言源代码开头都拥有一个package声明，表示源代码文件所属代码包

要生成Go语言可执行程序，必须要有main的package包，且必须在该包下有main()函数

同一个路径下只能存在一个package，一个package可以拆成多个源文件组成

# import

导入源代码文件所依赖的package包
不得导入源代码文件中没有用到的package，否则Go语言编译器会报编译错误
import语法
第一种：

    import "package1"
    import "package2"

第二种：

    import (
        "package1"
        "package2"
    )

# import原理

如果一个main导入其他包，包将被顺序导入
如果导入的包中依赖其他包（包B），会先导入B包，然后初始化B包中的常量和变量，最后如果B包中有init，会自动执行init()
所有包导入完成后才会对main中常量和变量进行初始化，然后执行main()中的init()函数（如果存在），最后执行main()函数
如果一个包被导入多次，则该包只会被导入一次
  

# import别名

别名操作：将导入的包命名为另一个容易记住的别名
例如： import format "fmt"

点（.）操作：点标识的包导入后，调用该包中函数时可以省略前缀包名
例如： import . "fmt"

下划线（_）操作：导入该包，但不导入整个包，而是执行该包中的init函数，因此无法通过包名来调用包中的其他函数，使用下划线操作往往是为了注册包里的引擎，让外部可以方便的使用
例如： import _ "fmt"


# 数据类型

数据类型的出现，是为了把数据分成所需内存大小不同的数据，编程的时候需腰用大数据的时候才需要申请大内存，就可以充分利用内存

数值类型
整型：uint8(byte), uint16, uint(uint32, uint64), int8, int16, int(int32(rune), int64) 一个字节8个位
浮点型：float32， float64
复数：complex64, complex128
无符号整型，存放指针uintptr
字符串类型，统一编码utf-8
布尔类型 常量true false

派生类型：
指针类型（Pointer）
数组类型
结构化类型（struct）
Channel类型（chan）
函数类型（func）
切片类型（slice）
接口类型（interface）
Map类型（map）

类型零值
类型零值不是控值，而是某个变量被声明后的默认值，一般情况下，值类型默认值为0，布尔类型默认值为false，string默认值为空字符串

类型查看
reflact.TypeOf(var)

类型别名
type a int
a 就是int的别名


类型所占存储大小
unsafe.Sizeof(var)


# 变量和常量

变量声明，初始化与赋值
1）变量声明格式：var <变量名称> [变量类型]
2）变量的赋值格式：<变量名称> = <值，表达式，函数>
3）声明和赋值同事进行：var <变量名称> [变量类型] = <值>
4) 分组声明：

    var(
        i int
        j float
        s string
    )

说明：
1)同一行声明多个变量和赋值

    var a, b, c int = 1, 2, 3（int 可以省略，自动进行类型推断） 

或者

    a, b := 1, 2 (只能用于局部变量)

2)全局变量的声明必须使用var关键字，局部变量则可以省略

3）特殊变量下划线“_” 相当于扔掉

# 变量类型转换

1）Go中不存在隐式转换，类型转换必须是显示的
2）类型转换只能发生在两种兼容类型之间
3）类型转换格式：<变量名称> [:] = <目标类型>(<需要转换的变量>)

# 变量可见性规则

大写字母开头的变量可以是导出的，也就是其他包可以读取的，是公用变量
小写字母开头的就是不可导出的，是私有变量

# 常量

定义形式
显示 `const identifier [type] = value`
隐式 `const identifier = value (无类型常量)`
常量可以使用内置表达式定义
如：len(), unsafe.Sizeof()

类型范围
目前支持布尔类型，数字类型（整数，浮点数，复数）和字符串

# 特殊常量iota

iota在const关键字出现时将被重置为0
const中每新增一行常量声明将使iota计数一次

iota常见用法
- 跳值使用法
- 插队使用法
- 表达式隐式使用
- 单行使用法

# 运算符

算术运算符
+, -, *, /, % ++, --(只能在变量后面用)

关系运算符
==, !=, >, <, >=, <=

逻辑运算符
&&, ||, !

按位运算符（一个字节8个位）
&, |, ^异或, <<左移（相当于*2）, >>右移

赋值运算符
=, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |=

# 控制流语句

条件语句if, if...else, 嵌套if, else if
选择语句switch, select
循环语句for
控制语句中使用到额关键字goto, break, continue


