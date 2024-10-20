package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 读取文件保存成str数组
func readFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func diffLines(lines1, lines2 []string) {
	m := make(map[string]int)
	for _, line := range lines1 {
		m[line]++
	}
	fmt.Println("文件2有文件1没有：")
	for _, line := range lines2 {
		if m[line] == 0 {
			fmt.Println("+", line)
		} else {
			m[line]--
		}
	}
	fmt.Println("文件1有文件2没有：")
	for line, count := range m {
		if count > 0 {
			fmt.Println("-", line)
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage：%s <fiile1><fiile1>", os.Args[0])
	}
	file1 := os.Args[1]
	file2 := os.Args[2]
	lines1, err := readFileLines(file1)
	if err != nil {
		log.Fatal(err)
	}
	lines2, err := readFileLines(file2)
	if err != nil {
		log.Fatal(err)
	}
	diffLines(lines1, lines2)
}
