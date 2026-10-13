package courier

import (
	"context"

	pb "github.com/feanchy/Courier-Service/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCourierServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) AssignCourier(
	ctx context.Context,
	req *pb.AssignOrderRequest,
) (*pb.AssignOrderResponse, error) {
	if req == nil || req.GetOrderId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "order_id must be positive")
	}

	courierID := int64(1)

	return &pb.AssignOrderResponse{
		CourierId: courierID,
	}, nil
}
