package main

import (
	"context"
	pb "gin-init/core/rpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func invoke(x, y int32) (sum, subtract int32) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//
	client := pb.NewCalculatorClient(conn)

	// 测试 Add 方法
	r, err := client.Add(ctx, &pb.AddRequest{A: x, B: y})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Add: %d", r.Result)
	sum = r.Result

	// 测试 Subtract 方法
	r2, err := client.Subtract(ctx, &pb.SubtractRequest{A: x, B: y})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Subtract: %d", r2.Result)
	subtract = r2.Result

	return
}

func invoke2(num int32) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//
	unaryStreamClient := pb.NewClientUnaryServerStreamClient(conn)

	// 测试 PrimeFactors 方法 (Unary -> Stream)
	stream, err := unaryStreamClient.PrimeFactors(ctx, &pb.PrimeFactorsRequest{Number: num})
	if err != nil {
		log.Fatalf("could not get prime factors: %v", err)
	}
	for {
		factor, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive a prime factor: %v", err)
		}
		log.Printf("Prime Factor ----> %d", factor.Factor)
	}

}

func invoke3(arr []int32) float32 {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//
	streamUnaryClient := pb.NewClientStreamServerUnaryClient(conn)

	// 测试 Average 方法 (Stream -> Unary)
	avgStream, err := streamUnaryClient.Average(ctx)
	if err != nil {
		log.Fatalf("could not calculate average: %v", err)
	}
	for _, number := range arr {
		if err := avgStream.Send(&pb.AverageRequest{Number: number}); err != nil {
			log.Fatalf("failed to send a number: %v", err)
		}
	}
	avgResponse, err := avgStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive the average: %v", err)
	}
	log.Printf("Average::::::: %f", avgResponse.Average)
	return avgResponse.Average
}

func invoke4() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//
	streamStreamClient := pb.NewClientStreamServerStreamClient(conn)

	// 测试 Chat 方法 (Stream -> Stream)
	chatStream, err := streamStreamClient.Chat(ctx)
	if err != nil {
		log.Fatalf("could not chat: %v", err)
	}
	messages := []string{"Hello", "How are you?", "Bye"}
	for _, message := range messages {
		if err := chatStream.Send(&pb.ChatMessage{Message: message}); err != nil {
			log.Fatalf("failed to send a message: %v", err)
		}
		reply, err := chatStream.Recv()
		if err != nil {
			log.Fatalf("failed to receive a reply: %v", err)
		}
		log.Printf("Received: %s", reply.Message)
	}
	chatStream.CloseSend()
}
