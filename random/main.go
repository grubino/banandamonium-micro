package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	proto "proto"
)

type Random struct{};

func generateRandomSequence(count int32, size int32) []int32 {
	seq := make([]int32, count)
	for i := range seq {
		seq[i] = rand.Int31n(size)
	}
	return seq
}

func (r *Random) RollDice(ctx context.Context, req *proto.RollRequest, rsp *proto.RollResponse) error {
	rsp.Results = append(rsp.Results, generateRandomSequence(2, 6)...)
	return nil
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	service := micro.NewService(
		micro.Name("random"),
	)
	service.Init()
	proto.RegisterBananaRandomHandler(service.Server(), new(Random))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
