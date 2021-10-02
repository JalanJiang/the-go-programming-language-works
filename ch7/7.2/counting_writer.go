package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	Writer io.Writer
	Count  int64
}

// 实现字节数计算
func (c *CountWriter) Write(content []byte) (int, error) {
	c.Count = int64(len(content))
	n, err := c.Writer.Write(content)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{
		Writer: w,
	}
	return &cw, &cw.Count
}

func main() {
	rw, ilen := CountingWriter(os.Stdout)
	rw.Write([]byte("hello,world\n"))
	fmt.Println(*ilen)
}
