package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

// 计算单词行数
func (w *WordCounter) Write(p []byte) (int, error) {
	// 通过 NewReader 创建 bufio.Reader 对象
	r := bytes.NewReader(p)
	// 通过 NewScanner 创建 Scanner 对象
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w++
	}

	return len(p), nil
}

type LineCounter int

// 行数计算
func (l *LineCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l++
	}

	return len(p), nil
}

func main() {
	// 单词计算
	var w WordCounter
	wordString := "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"
	w.Write([]byte(wordString))
	fmt.Println(w)
	w = 0
	fmt.Fprintf(&w, "hello word %s", wordString)
	fmt.Println(w)

	// 行数计算
	var l LineCounter
	lineString := "hello word!\nMy name is Jalan!\nHi!\nHahaha!"
	l.Write([]byte(lineString))
	fmt.Println(l)
	l = 0
	fmt.Fprintf(&l, "Add line\n%s", lineString)
	fmt.Println(l)
}
