package server

import (
	"context"
	"fmt"
	"gin-init/config"
	pb "gin-init/core/rpc/proto"
	"google.golang.org/grpc"
	"io"
	"net"
)

type RPCServer struct {
	pb.UnimplementedCalculatorServer
	pb.UnimplementedClientUnaryServerStreamServer
	pb.UnimplementedClientStreamServerUnaryServer
	pb.UnimplementedClientStreamServerStreamServer
}

func (s *RPCServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	fmt.Printf("receviced ----> %v", req)
	return &pb.AddResponse{Result: req.A + req.B}, nil
}

func (s *RPCServer) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	fmt.Printf("receviced ---->  %v", req)
	return &pb.SubtractResponse{Result: req.A - req.B}, nil
}

// 实现 ClientUnaryServerStream 服务
func (s *RPCServer) PrimeFactors(req *pb.PrimeFactorsRequest, stream pb.ClientUnaryServerStream_PrimeFactorsServer) error {
	fmt.Printf("Received PrimeFactors request: %v\n", req)
	number := req.Number
	factor := int32(2)
	for number > 1 {
		if number%factor == 0 {
			stream.Send(&pb.PrimeFactorsResponse{Factor: factor})
			number /= factor
		} else {
			factor++
		}
	}
	return nil
}

// 实现 ClientStreamServerUnary 服务
func (s *RPCServer) Average(stream pb.ClientStreamServerUnary_AverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		fmt.Printf("receviced ----> %v", req)
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{Average: float32(sum) / float32(count)})
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}
}

// 实现 ClientStreamServerStream 服务
func (s *RPCServer) Chat(stream pb.ClientStreamServerStream_ChatServer) error {
	for {
		msg, err := stream.Recv()
		fmt.Println("msg  ---->  ", msg)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received message: %v\n", msg.Message)
		err = stream.Send(&pb.ChatMessage{Message: "Echo: " + msg.Message})
		if err != nil {
			return err
		}
	}
}

func StartGPRC() {
	listen, _ := net.Listen("tcp", config.Conf.GRPC.GetAddr())
	// 创建grpc服务
	newServer := grpc.NewServer()
	pb.RegisterCalculatorServer(newServer, &RPCServer{})
	pb.RegisterClientUnaryServerStreamServer(newServer, &RPCServer{})
	pb.RegisterClientStreamServerUnaryServer(newServer, &RPCServer{})
	pb.RegisterClientStreamServerStreamServer(newServer, &RPCServer{})

	// 启动服务
	go func() {
		err := newServer.Serve(listen)
		if err != nil {
			fmt.Println("[Error] 启动grpc服务器失败！ ", err.Error())
		}
		fmt.Println("gRPC 服务启动成功 in ", config.Conf.GRPC.GetAddr())
	}()
}
