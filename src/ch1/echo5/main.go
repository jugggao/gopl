// 修改 echo 程序，使其打印每个参数的索引和值，每个一行
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", index, arg)
	}
}
