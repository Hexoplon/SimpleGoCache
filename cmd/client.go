package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/Hexoplon/SimpleGoCache"
	"github.com/micro/go-micro"
)

// A test client to verify functionality
func main() {
	service := micro.NewService(micro.Name("LobsterCache.client"))
	service.Init()

	cache := proto.NewCacheService("LobsterCache", service.Client())

	rsp, err := cache.NewCache(context.TODO(), &proto.NewCacheMsg{
		Key:            "test",
		Ttl:            120,
		MaxElements:    10,
		MaxElementSize: 4 * proto.KB,
	})
	if err != nil {
		log.Println("Error:", err)
	}
	if rsp != nil {
		fmt.Println(rsp.Msg)
	}
}
