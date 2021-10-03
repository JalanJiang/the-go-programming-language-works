package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 创建一个通道
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 遍历参数，开协程请求
		go fetchTwoTimes(url, ch)
	}

	for range os.Args[1:] {
		// 从通道中取出结果
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchTwoTimes(url string, ch chan string) {
	ch <- "url: " + url + "\nresp1: " + fetch(url) + "\n" + "resp2: " + fetch(url)
}

func fetch(url string) string {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprint(err)
	}

	// 写到 ioutil.Discard 输出流进行丢弃
	io.Copy(os.Stdout, resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	// 写入到文件
	// out, err := os.Create("/out.txt")
	// wt := bufio.NewWriter(out)
	// defer out.Close()
	// // 按一定大小循环写入数据，避免大文件造成内存溢出
	// io.Copy(wt, resp.Body)
	// wt.Flush()

	if err != nil {
		return fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// go run main.go https://www.baidu.com https://www.sina.com
