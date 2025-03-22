package grpc

import (
	"context"
	pb "github.com/TemaStatham/TaskService/proto/profile"
	Organization "github.com/TemaStatham/TaskService/taskservice/pkg/app/organization/model"
	organizationquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/organization/query"
	User "github.com/TemaStatham/TaskService/taskservice/pkg/app/user/model"
	userquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/user/query"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ClientInterface interface {
	Close()
	userquery.ClientUserInterface
	organizationquery.ClientOrganizationInterface
}

type Client struct {
	Client pb.ProfileServiceClient
	Conn   *grpc.ClientConn
}

func NewGrpcClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: pb.NewProfileServiceClient(conn),
		Conn:   conn,
	}, nil
}

func (c *Client) Close() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}

func (c *Client) GetUser(ctx context.Context, userID uint64) (User.User, error) {
	res, err := c.Client.GetUser(ctx, &pb.UserRequest{Id: userID})
	if err != nil {
		return User.User{}, err
	}

	log.Printf("Пользователь: %s %s, Админ: %v", res.Name, res.Surname, res.IsAdmin)
	return User.User{
		Name:    res.Name,
		Surname: &res.Surname,
		IsAdmin: res.IsAdmin,
	}, nil
}

func (c *Client) GetOrganization(ctx context.Context, orgID uint64) (Organization.Organization, error) {
	res, err := c.Client.GetOrganization(ctx, &pb.OrganizationRequest{Id: orgID})
	if err != nil {
		return Organization.Organization{}, err
	}

	log.Printf("Организация: Email - %s, StatusID - %d", res.Email, res.StatusId)
	return Organization.Organization{
		StatusID: uint(res.StatusId),
		ID:       uint(orgID),
		Name:     res.Email,
	}, nil
}

func (c *Client) GetOrganizationsByUserID(ctx context.Context, userID uint64) ([]Organization.Organization, error) {
	res, err := c.Client.GetOrganizationsByUserID(ctx, &pb.OrganizationUserRequest{Id: userID})
	if err != nil {
		return []Organization.Organization{}, err
	}

	var orgs []Organization.Organization
	log.Println("Организации пользователя:")
	for _, org := range res.Organizations {
		log.Printf("ID: %d, Владелец: %v", org.Id, org.IsOwner)
		orgs = append(orgs, Organization.Organization{
			ID: uint(org.Id),
		})
	}

	return orgs, nil
}

func (c *Client) GetUsersByIDS(ctx context.Context, users []uint) ([]User.User, error) {
	res, err := c.Client.GetUsersByIDS(ctx, &pb.GetUsersByIDsRequest{})
	if err != nil {
		return []User.User{}, err
	}

	usersres := []User.User{}
	for _, user := range res.Users {
		usersres = append(usersres, User.User{
			ID: uint(user.Id),
		})
	}

	return usersres, nil
}
