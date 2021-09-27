package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ittus/grpc-go/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()

	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}

func (*server) PrimeNumber(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberServer) error {
	fmt.Printf("PrimeNumber function was invoked with %v", req)
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send((&calculatorpb.PrimeNumberResponse{
				PrimeFactor: int64(divisor),
			}))
			number = number / divisor
		} else {
			divisor++
			fmt.Println("Divisor has increased to %v", divisor)
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
