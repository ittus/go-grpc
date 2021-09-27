package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ittus/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)
	// doUnary(c)

	// doServerStreaming(c)

	doClientStream(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 3,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator GPC: %v", err)
	}
	log.Printf("Reponse from Calculator: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Streaming RPC...")
	req := &calculatorpb.PrimeNumberRequest{
		Number: 12,
	}

	stream, err := c.PrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator GPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStream(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do ComputeAverage ClientStreaming RPC...")
	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{1, 2, 3, 4, 10}
	for _, number := range numbers {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reading response from stream: %v", err)
	}
	fmt.Printf("The average is: %v", res.GetAverage())
}
