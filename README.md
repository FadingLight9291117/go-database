# copy of SQLite Database System

[copy of SQLite Database System: Design and Implementation](https://cstack.github.io/db_tutorial/parts/part1.html)

- [copy of SQLite Database System](#copy-of-sqlite-database-system)
    - [标准`IO`](#标准io)
        - [标准输入](#标准输入)
        - [标准输出](#标准输出)
    - [文件路径操作](#文件路径操作)
    - [go内存分配 `tcmalloc`](#go内存分配-tcmalloc)
    - [string 和 rune 的一些转换](#string-和-rune-的一些转换)
    - [`struct` 和 `[]byte` 的相互转化](#struct-和-byte-的相互转化)
    - [some problem](#some-problem)
        - [2022.8.1](#202281)

## 标准`IO`

> Go 的标准输入，如何读取一行输入

因为 `fmt.Scan` `fmt.Scanf` `fmt.Scanln` 都是遇到空格结束输入，所以都无法读取一行输入

有两种解决方法

### 标准输入

```go
// 1. bufio.NewScanner
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
input := scanner.Text()
input = string.TrimSpace(input)

// 2. bufio.NewReader
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReaderString('\n')
input = strings.TrimSpace(input)

```

### 标准输出

```go
fmt.Printf
fmt.Print
fmt.Println

// 以下输出到标准错误流
print
println
```

## 文件路径操作

```go


```

## go内存分配 `tcmalloc`

> 内存分配器

[tcmalloc 介绍 | Legendtkl](http://legendtkl.com/2015/12/11/go-memory/)

## string 和 rune 的一些转换

```go

// `string` to `[]rune`
r1 := []rune([string])

// `[]rune` to `[...]rune` 切片转数组
r2 := [32]rune{}
copy(r2[:], r1)

// `[]rune` to `string`
string(r1)

// `[...]rune` to `string`
string(r1[:])

```

## `struct` 和 `[]byte` 的相互转化

> Data must be a fixed-size value or a slice of fixed-size values, or a pointer to such data. Boolean values encode as
> one byte: 1 for true, and 0 for false. Bytes written to w are encoded using the specified byte order and read from
> successive fields of the data. When writing structs, zero values are written for fields with blank (_) field names.
>
> 数据必须是

## some problem

### 2022.8.1

1. `string`不能直接写入 `bytes.Buffer`, 需要转换为`rune array`

```go
buf := new(bytes.Buffer)
binary.Write(buf, b)
```

### 2022.8.2

1. `for range`遍历切片或者数组时注意

```go
for i, v := range rows {
&v // 这里的`v`的地址一直都是同一个		
// 因为每次遍历都是将值拷贝给`v` 
// 而`v`一直被复用，所以地址一直不会变
}
```

2. `const`

常量是在编译时被创建的，即使定义在函数内部也是如此

只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。

### 2022.8.23

1. 从文件读入后再次插入的错误原因是，每个Page的长度计算错误，原来是`PageSize`，是`row`的个数，应该乘上`RowSize`

2. 第二个错误是`File.ReadAt`这个函数，传入的`[]byte`长度不能大于`File`文件的长度；

### 2022.9.17

1. `unsafe.SizeOf`函数的问题

```go
type Foo2 struct {
Test      uint8
ParentPtr int64
}

func main() {
foo := Foo2{}
fmt.Printf("CommonNodeHeader size is %d\n", unsafe.Sizeof(foo))

buf := &bytes.Buffer{}
err := binary.Write(buf, binary.BigEndian, foo)
if err != nil {
return
}

fmt.Printf("CommonNodeHeader byes size of is %d\n", len(buf.Bytes()))
}
```

输出

```go
CommonNodeHeader size is 16
CommonNodeHeader byes size of is 9
```

明显结构体的SizeOf不正确

原因：结构体的字节对齐

https://www.jianshu.com/p/a2e157c95d9e

> “Go是一个C家族语言，Go的结构体类型基于C语言的结构体演化而来，因此关于字节对齐等概念也是通用的。”

 https://cloud.tencent.com/developer/article/1817214
 

> - 在所有结构体成员的字节长度都没有超出操作系统基本字节单位(32位操作系统是4,64位操作系统是8)的情况下，按照结构体中字节最大的变量长度来对齐；
> - 若结构体中某个变量字节超出操作系统基本字节单位，那么就按照系统字节单位来对齐。

> 字节对齐的根本原因其实在于cpu读取内存的效率问题，对齐以后，cpu读取内存的效率会更快。但是这里有个问题，就是对齐的时候0x00000002~0x00000004
> 这三个字节是浪费的，所以字节对齐实际上也有那么点以空间换时间的意思，具体写代码的时候怎么选择，其实是看个人的。



## 扫描多层的B树

为了在到达第一个叶子节点的末尾时，跳转到第二个叶子节点，需要在叶子节点的头部增加一个字段 `NextLeaf`保存右侧兄弟节点的页码