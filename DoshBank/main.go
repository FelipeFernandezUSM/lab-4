package main

import (
    "context"
    "log"
    "net"

    pb "github.com/FelipeFernandezUSM/lab-4/Comunicacion"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

type server struct{}

func (s *server) RequestMoney(ctx context.Context, req *comunicacion.MoneyRequestToDoshbank) (*comunicacion.MoneyResponseFromDoshbank, error) {
    log.Printf("Received money request from Director: %s", req.Message)

    // Simulate providing the requested amount of money
    amount := int32(100) // You can replace this with your logic to provide the money

    return &comunicacion.MoneyResponseFromDoshbank{Amount: amount}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052") // Use a different port than Director and Jugador
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    comunicacion.RegisterDoshbankServiceServer(grpcServer, &server{})
    log.Println("Doshbank Server listening at :50052")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}