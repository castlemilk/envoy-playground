package calculator

import (
	"context"
	"fmt"
	"log"
	pb "service/protos/calculator"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Calculate(ctx context.Context, r *pb.BinaryOperation) (*pb.CalculationResult, error) {
	log.Println("[server:Calculate] Started")
	if ctx.Err() == context.Canceled {
		return &pb.CalculationResult{}, fmt.Errorf("client cancelled: abandoning")
	}

	switch r.GetOperation() {
	case pb.Operation_ADD:
		return &pb.CalculationResult{
			Result: r.GetFirstOperand() + r.GetSecondOperand(),
		}, nil
	case pb.Operation_SUBTRACT:
		return &pb.CalculationResult{
			Result: r.GetFirstOperand() - r.GetSecondOperand(),
		}, nil
	default:
		return &pb.CalculationResult{}, fmt.Errorf("undefined operation")
	}
}
