// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//数组定义：var numbers [5]int
	//numbers := [5]int{1,2,3,4,5}
	//m := make(map[int]string, 2)
	//map只能传入2个参数，不能传入3个参数，切片可以有个容量参数
	counts := make(map[string]int)
	//Args：命令行参数，程序启动后输入的参数
	//files := os.Args[1:]
	//0代表命令行启动的那个exe文件的全路径
	strings := os.Args[0:]
	fmt.Println("000000", strings)
	dir, _ := os.Getwd()
	files := []string{dir + "/ch1/dup2/file1.txt"}
	fmt.Println("-------:", dir)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
