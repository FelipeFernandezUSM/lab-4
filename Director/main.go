package main

import (
    "context"
    "log"
    "net"

    pb "github.com/FelipeFernandezUSM/lab-4/Comunicacion"
    "google.golang.org/grpc"
)

type server struct{}

func (s *server) SendActNow(ctx context.Context, in *comunicacion.ActNow) (*emptypb.Empty, error) {
    log.Printf("Received ActNow: %v", in.GetActNow())
    return &emptypb.Empty{}, nil
}

func (s *server) SendPlayerAlive(ctx context.Context, in *comunicacion.PlayerAlive) (*emptypb.Empty, error) {
    log.Printf("Received PlayerAlive: %v", in.GetPlayerAlive())
    return &emptypb.Empty{}, nil
}

func (s *server) SendOptionMessage(ctx context.Context, in *comunicacion.OptionMessage) (*emptypb.Empty, error) {
    log.Printf("Received OptionMessage: %v", in.GetOption())
    return &emptypb.Empty{}, nil
}

func (s *server) SendLetterMessage(ctx context.Context, in *comunicacion.LetterMessage) (*emptypb.Empty, error) {
    log.Printf("Received LetterMessage: %v", in.GetLetter())
    return &emptypb.Empty{}, nil
}

func (s *server) SendIntStringMessage(ctx context.Context, in *comunicacion.IntStringMessage) (*emptypb.Empty, error) {
    log.Printf("Received IntStringMessage: %v", in.GetIntString())
    return &emptypb.Empty{}, nil
}

func (s *server) RequestMoney(ctx context.Context, in *comunicacion.MoneyRequest) (*comunicacion.MoneyResponse, error) {
    log.Printf("Received MoneyRequest from %v: %v", in.GetName(), in.GetMessage())
    // Example response
    return &comunicacion.MoneyResponse{Amount: 100}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    comunicacion.RegisterComunicacionServiceServer(grpcServer, &server{})
    log.Printf("Director Server listening at %v", lis.Addr())

    // Start the Director server
    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

    // Set up connection to the Doshbank server
    conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect to Doshbank server: %v", err)
    }
    defer conn.Close()
    doshbankClient := comunicacion.NewDoshbankServiceClient(conn)

    // Send request for money to Doshbank
    moneyResponse, err := doshbankClient.RequestMoney(context.Background(), &comunicacion.MoneyRequestToDoshbank{Message: "Please send money"})
    if err != nil {
        log.Fatalf("failed to request money from Doshbank: %v", err)
    }
    log.Printf("Received money from Doshbank: %d", moneyResponse.Amount)
}