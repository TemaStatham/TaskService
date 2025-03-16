package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/TemaStatham/TaskService/functions"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProfileServiceServer
}

// GetUser — метод для получения информации о пользователе
func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received GetUser request for user ID: %d", req.GetId())

	// Логика получения информации о пользователе (например, из базы данных)
	// Примерный ответ
	userResponse := &pb.UserResponse{
		Name:    "John",
		Surname: "Doe",
		IsAdmin: true,
	}

	return userResponse, nil
}

// GetOrganization — метод для получения информации об организации
func (s *server) GetOrganization(ctx context.Context, req *pb.OrganizationRequest) (*pb.OrganizationResponse, error) {
	log.Printf("Received GetOrganization request for organization ID: %d", req.GetId())

	// Логика получения информации об организации
	// Примерный ответ
	organizationResponse := &pb.OrganizationResponse{
		Email:    "contact@organization.com",
		StatusId: 1, // Статус может быть, например, "активен"
	}

	return organizationResponse, nil
}

// GetOrganizationsByUserID — метод для получения всех организаций пользователя
func (s *server) GetOrganizationsByUserID(ctx context.Context, req *pb.OrganizationUserRequest) (*pb.OrganizationUserListResponse, error) {
	log.Printf("Received GetOrganizationsByUserID request for user ID: %d", req.GetId())

	// Логика получения всех организаций, в которых состоит пользователь
	// Примерный ответ с несколькими организациями
	organizations := []*pb.OrganizationUserResponse{
		{
			Id:      1,
			IsOwner: true,
		},
		{
			Id:      2,
			IsOwner: false,
		},
	}

	return &pb.OrganizationUserListResponse{
		Organizations: organizations,
	}, nil
}

func main() {
	// Настройка порта
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "50501"
	}

	// Настройка TCP сервера
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error in Listen: %v", err)
	}

	// Настройка gRPC сервера
	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &server{})
	log.Printf("listening at: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
