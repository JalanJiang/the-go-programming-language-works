package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"test/9/9.8"
)

func incomingURLs() []string {
	urls := []string{
		"https://www.baidu.com",
		"https://www.sina.com",
		"http://jalan.space",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com"}
	return urls
}

// Func 函数
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	// 协程执行的函数是 httpBody
	m := memo.New(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, false)
		if err != nil {
			log.Print(err)
		}
		// fmt.Printf("%s\n", url)
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}
