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

> Data must be a fixed-size value or a slice of fixed-size values, or a pointer to such data. Boolean values encode as one byte: 1 for true, and 0 for false. Bytes written to w are encoded using the specified byte order and read from successive fields of the data. When writing structs, zero values are written for fields with blank (_) field names.
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
	&v  // 这里的`v`的地址一直都是同一个		
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