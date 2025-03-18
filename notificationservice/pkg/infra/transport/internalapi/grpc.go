package internalapi

import (
	"context"
	"fmt"
	pb "github.com/TemaStatham/TaskService/proto/notification"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnsafeNotificationServiceServer
}

func (s Server) SendNotification(ctx context.Context, request *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	// todo: получить уведолмения, сохранить, отправить

	return &pb.NotificationResponse{}, nil
}

func NewServer() *Server {
	return &Server{}
}

func InitServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error in Listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &Server{})
	log.Printf("listening at: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
