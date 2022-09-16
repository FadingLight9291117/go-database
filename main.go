package main

import (
	"bufio"
	"com.fadinglight/db/executor"
	"com.fadinglight/db/processor"
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"flag"
	"fmt"
	"os"
	"strings"
)

func printPrompt() {
	fmt.Print("db> ")
}

func readInput(scanner *bufio.Scanner) *types.InputBuffer {
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)

	return &types.InputBuffer{Buffer: input}
}

/** init
 * 1. 打开数据库文件;
 * 2. 初始化 `pager` 数据结构;
 * 3. 初始化 `table` 数据结构。
 */
func dbOpen(filename string) (*table.Table, bool) {
	t := table.NewTable(filename)

	return t, true
}

var filename = flag.String("filename", "db.db", "数据库文件路径")

func main() {
	flag.Parse()
	if *filename == "" {
		panic("")
	}
	t, ok := dbOpen(*filename)
	if !ok {
		print("db open error.")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// 提示信息
		printPrompt()
		// 读取输入
		inputBuffer := readInput(scanner)
		if inputBuffer.Buffer == "" {
			continue
		}
		// 处理 `meta-command`
		// 类似 `.exit`以 `.`开头的命令被称为 `meta-command`(元命令)
		if inputBuffer.Buffer[0] == '.' {
			_, err := executor.DoMetaCommand(inputBuffer, t)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		// 解析普通命令
		statement := &types.Statement{}
		_, err := processor.PrepareStatement(inputBuffer, statement)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 执行
		_, err = executor.ExecuteStatement(statement, t)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Executed.")
	}
}
