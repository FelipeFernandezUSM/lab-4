package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/FelipeFernandezUSM/lab-4/Comunicacion" // Update this import path
    "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
    pb.UnimplementedComunicacionServiceServer
}

func (s *server) SendActNow(ctx context.Context, in *pb.ActNow) (*emptypb.Empty, error) {
    log.Printf("Received ActNow: %v", in.GetActNow())
    return &emptypb.Empty{}, nil
}

func (s *server) SendPlayerAlive(ctx context.Context, in *pb.PlayerAlive) (*emptypb.Empty, error) {
    log.Printf("Received PlayerAlive: %v", in.GetPlayerAlive())
    return &emptypb.Empty{}, nil
}

func (s *server) SendOptionMessage(ctx context.Context, in *pb.OptionMessage) (*emptypb.Empty, error) {
    log.Printf("Received OptionMessage: %v", in.GetOption())
    return &emptypb.Empty{}, nil
}

func (s *server) SendLetterMessage(ctx context.Context, in *pb.LetterMessage) (*emptypb.Empty, error) {
    log.Printf("Received LetterMessage: %v", in.GetLetter())
    return &emptypb.Empty{}, nil
}

func (s *server) SendIntStringMessage(ctx context.Context, in *pb.IntStringMessage) (*emptypb.Empty, error) {
    log.Printf("Received IntStringMessage: %v", in.GetIntString())
    return &emptypb.Empty{}, nil
}

func (s *server) RequestMoney(ctx context.Context, in *pb.MoneyRequest) (*pb.MoneyResponse, error) {
    log.Printf("Received MoneyRequest from %v: %v", in.GetName(), in.GetMessage())
    // Example response
    return &pb.MoneyResponse{Amount: 100}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterComunicacionServiceServer(grpcServer, &server{})
    log.Printf("Server listening at %v", lis.Addr())
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}