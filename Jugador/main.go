package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "github.com/FelipeFernandezUSM/lab-4/Comunicacion" // Update this import path
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewComunicacionServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Sending ActNow message
    _, err = c.SendActNow(ctx, &pb.ActNow{ActNow: true})
    if err != nil {
        log.Fatalf("could not send ActNow: %v", err)
    }
    log.Println("Sent ActNow")

    // Sending PlayerAlive message
    _, err = c.SendPlayerAlive(ctx, &pb.PlayerAlive{PlayerAlive: true})
    if err != nil {
        log.Fatalf("could not send PlayerAlive: %v", err)
    }
    log.Println("Sent PlayerAlive")

    // Sending OptionMessage
    _, err = c.SendOptionMessage(ctx, &pb.OptionMessage{Option: 1})
    if err != nil {
        log.Fatalf("could not send OptionMessage: %v", err)
    }
    log.Println("Sent OptionMessage")

    // Sending LetterMessage
    _, err = c.SendLetterMessage(ctx, &pb.LetterMessage{Letter: "A"})
    if err != nil {
        log.Fatalf("could not send LetterMessage: %v", err)
    }
    log.Println("Sent LetterMessage")

    // Sending IntStringMessage
    _, err = c.SendIntStringMessage(ctx, &pb.IntStringMessage{IntString: "1,2,3,4,5"})
    if err != nil {
        log.Fatalf("could not send IntStringMessage: %v", err)
    }
    log.Println("Sent IntStringMessage")

    // Requesting Money
    res, err := c.RequestMoney(ctx, &pb.MoneyRequest{Name: "Jugador1", Message: "Please send money"})
    if err != nil {
        log.Fatalf("could not request money: %v", err)
    }
    log.Printf("Received MoneyResponse: %v", res.GetAmount())
}