package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	// 连接 es 客户端（本机且无帐号密码）
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
}
