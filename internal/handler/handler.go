package handler

import (
	"context"
	"github.com/TemaStatham/TaskService/internal/handler/request"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/hub"
	"github.com/TemaStatham/TaskService/pkg/middleware/auth"
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	InvalidInputBodyErr = "invalid input body"
)

type TaskService interface {
	Get(ctx context.Context, id uint, user uint) (*model.Task, error)
	Show(ctx context.Context, paginat *paginate.Pagination, user uint) (*paginate.Pagination, error)
	Create(ctx context.Context, dto request.CreateTaskRequest, user uint) (uint, error)
	Update(ctx context.Context, dto request.UpdateTaskRequest, user uint) error
	Delete(ctx context.Context, id uint, user uint) error
}

type ResponseService interface {
	Create(ctx context.Context, dto request.CreateResponseRequest, user uint) (uint, error)
	Show(ctx context.Context, dto request.GetResponseRequest, user uint) (*paginate.Pagination, error)
	Update(ctx context.Context, dto request.UpdateResponseRequest, user uint) error
}

type CommentService interface {
	Create(ctx context.Context, dto request.CreateCommentRequest, user uint) (uint, error)
	Show(ctx context.Context, dto request.ShowCommentRequest, user uint) (*paginate.Pagination, error)
}

type Handler struct {
	TaskService
	ResponseService
	CommentService
}

func NewTaskHandler(taskService TaskService, responseService ResponseService, commentService CommentService) *Handler {
	return &Handler{
		taskService,
		responseService,
		commentService,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(auth.UserIdentity())

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

	wsHub := hub.NewHub()
	go wsHub.Run()

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		hub.ServeWS(c, roomId, wsHub)
	})

	return router
}
