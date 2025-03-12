package handler

import (
	"context"
	hub2 "github.com/TemaStatham/TaskService/internal/app/hub"
	"github.com/TemaStatham/TaskService/internal/handler/request"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/middleware/auth"
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	InvalidInputBodyErr = "invalid input body"
)

type TaskService interface {
	Get(ctx context.Context, id uint) (*model.TaskModel, error)
	Update(ctx context.Context, dto *request.UpdateTaskRequest) error
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, dto *request.CreateTaskRequest) (uint, error)
	Show(
		ctx context.Context,
		pagination *paginate.Pagination,
		user uint,
	) (*paginate.Pagination, error)
}

type ResponseService interface {
	Create(
		ctx context.Context,
		dto *request.CreateResponseRequest,
		user uint,
	) (uint, error)
	Show(
		ctx context.Context,
		dto *request.GetResponseRequest,
	) (*paginate.Pagination, error)
	Update(
		ctx context.Context,
		dto *request.UpdateResponseRequest,
	) error
}

type CommentService interface {
	Create(
		ctx context.Context,
		dto *request.CreateCommentRequest,
		user uint,
	) (uint, error)
	Show(
		ctx context.Context,
		dto *request.ShowCommentRequest,
	) (*paginate.Pagination, error)
}

type Handler struct {
	TaskService
	ResponseService
	CommentService
}

func NewTaskHandler(
	taskService TaskService,
	responseService ResponseService,
	commentService CommentService,
) *Handler {
	return &Handler{
		taskService,
		responseService,
		commentService,
	}
}

func (h *Handler) Init(jwtSecret string) *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(auth.UserIdentity(jwtSecret))

	tasks := router.Group("/tasks")
	{
		tasks.GET("/", h.getTasks)
		tasks.GET("/:id", h.getTask)
		tasks.POST("/", h.createTask)
		tasks.PUT("/:id", h.updateTask)
		tasks.DELETE("/:id", h.deleteTask)
	}

	responses := router.Group("/responses")
	{
		responses.GET("/", h.getResponses)
		responses.POST("/", h.createResponse)
		responses.PUT("/:id", h.updateResponse)
	}

	wsHub := hub2.NewHub()
	go wsHub.Run()

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		hub2.ServeWS(c, roomId, wsHub)
	})

	return router
}
