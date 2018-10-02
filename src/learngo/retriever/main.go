package main

import (
	"fmt"
	"time"

	"./mock"
	"./real"
)

type Retriever interface {
	Get(url string) string
}

func dowload(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fake retriever"}
	r = real.Retriever{UserAgent: "chrom", TimeOut: time.Minute}
	fmt.Println(dowload(r))
}
