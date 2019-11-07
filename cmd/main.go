package main

import (
	"fmt"
	"log"

	proto "github.com/Hexoplon/SimpleGoCache"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("LobsterCache"))

	service.Init()

	err := proto.RegisterCacheHandler(service.Server(), &proto.CacheStore{})
	if err != nil {
		log.Fatalln(proto.ErrorStartCacheHandler, err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
