package courier

import (
	"context"
	"testing"

	pb "github.com/feanchy/Courier-Service/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAssignCourier_Success(t *testing.T) {
	server := NewServer()

	resp, err := server.AssignCourier(context.Background(), &pb.AssignOrderRequest{
		OrderId: 123,
	})
	if err != nil {
		t.Fatalf("AssignCourier() error = %v", err)
	}

	if resp.GetCourierId() != 1 {

		t.Fatalf("AssignCourier() courier_id = %d, want 1", resp.GetCourierId())
	}

}

func TestAssignCourier_InvalidOrderID(t *testing.T) {
	server := NewServer()

	_, err := server.AssignCourier(context.Background(), &pb.AssignOrderRequest{
		OrderId: 0,
	})
	if err == nil {
		t.Fatal("AssignCourier expected error, got nil")
	}

	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("AssignCourier() error code = %s, want %s",
			status.Code(err), codes.InvalidArgument)
	}
}
