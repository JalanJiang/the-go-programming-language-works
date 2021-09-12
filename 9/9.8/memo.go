package memo

import (
	"io/ioutil"
	"net/http"
)

// 放置结果
type entry struct {
	res   result
	ready chan struct{} // 结果是否准备好
}

// 请求消息
type request struct {
	key      string
	done     chan bool     // done 通道，用于取消缓存
	response chan<- result // 响应通道
}

// Memo 缓存了调用 Func 的结果
type Memo struct {
	// 一个用于传输 request 结构的通道
	requests chan request
}

// 用于记忆的函数类型
type Func func(key string) (interface{}, error)

// 请求结果
type result struct {
	value interface{}
	err   error
}

// 用协程调用 f
func New(f Func) *Memo {
	// 包含用于传输 request 的通道
	memo := &Memo{requests: make(chan request)}
	// 用协程执行 server，f 的参数是 url
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, isDone bool) (interface{}, error) {
	// 创建一个结果通道
	response := make(chan result)
	// 作业：
	// 创建一个 done 通道，用于取消操作
	// 不缓存取消操作的结果
	done := make(chan bool)
	// 向 requests 通道放入结构体，需要处理的 request
	request := request{key, done, response}
	memo.requests <- request
	// 需要放在这里发送数据，否则被阻塞无法处理
	request.done <- isDone
	// 取出结果
	res := <-response

	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	// 创建一个 cache
	cache := make(map[string]*entry)
	// 遍历通道里的数据
	for req := range memo.requests {
		// 缓存里是否有数据
		e := cache[req.key]
		// 判断是否需要缓存
		isDone := <-req.done
		// isDone := true
		if e == nil || isDone == false {
			// 没有数据，重新创建
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			// 用协程执行 call，call 执行了 f 函数
			go e.call(f, req.key)
		}
		// 分发数据：当数据准备完毕，向 response 通道传输结果
		go e.deliver(req.response)
	}
}

// 执行 f 后关闭通道
func (e *entry) call(f Func, key string) {
	// 执行函数
	e.res.value, e.res.err = f(key)
	// 通知数据已准备完毕
	close(e.ready)
}

// 分发数据
func (e *entry) deliver(response chan<- result) {
	// 等待数据准备完毕
	<-e.ready
	// 向客户端发送结果
	response <- e.res
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

// func main() {
// 	// 协程执行的函数是 httpBody
// 	m := New(httpGetBody)
// 	for _, url := range incomingURLs() {
// 		start := time.Now()
// 		value, err := m.Get(url, false)
// 		if err != nil {
// 			log.Print(err)
// 		}
// 		// fmt.Printf("%s\n", url)
// 		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
// 	}
// }
