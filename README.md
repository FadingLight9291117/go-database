# copy of SQLite Database System

[copy of SQLite Database System: Design and Implementation](https://cstack.github.io/db_tutorial/parts/part1.html)

## 目录

## `tcmalloc`

> 内存分配器

[tcmalloc 介绍 | Legendtkl](http://legendtkl.com/2015/12/11/go-memory/)


## string和rune的一些转换

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