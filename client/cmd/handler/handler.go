package handler

import (
	"fmt"
	"github.com/TemaStatham/TaskService/client/cmd/handler/hub"
	approveservice "github.com/TemaStatham/TaskService/client/pkg/app/approve/service"
	commentquery "github.com/TemaStatham/TaskService/client/pkg/app/comment/query"
	commentservice "github.com/TemaStatham/TaskService/client/pkg/app/comment/service"
	responsequery "github.com/TemaStatham/TaskService/client/pkg/app/response/query"
	responseservice "github.com/TemaStatham/TaskService/client/pkg/app/response/service"
	taskquery "github.com/TemaStatham/TaskService/client/pkg/app/task/query"
	taskservice "github.com/TemaStatham/TaskService/client/pkg/app/task/service"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/middleware/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

const (
	InvalidInputBodyErr = "invalid input body"
)

type Handler struct {
	responseQuery   responsequery.ResponseQueryInterface
	responseService responseservice.ResponseServiceInterface
	commentQuery    commentquery.CommentQueryInterface
	commentService  commentservice.CommentServiceInterface
	taskQuery       taskquery.TaskQueryInterface
	taskService     taskservice.TaskServiceInterface
	approveService  approveservice.ApproveServiceInterface
}

func NewTaskHandler(
	responseQuery responsequery.ResponseQueryInterface,
	responseService responseservice.ResponseServiceInterface,
	commentQuery commentquery.CommentQueryInterface,
	commentService commentservice.CommentServiceInterface,
	taskQuery taskquery.TaskQueryInterface,
	taskService taskservice.TaskServiceInterface,
	approveService approveservice.ApproveServiceInterface,
) *Handler {
	return &Handler{
		responseQuery,
		responseService,
		commentQuery,
		commentService,
		taskQuery,
		taskService,
		approveService,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(ctx *gin.Context, roomID, userID uint, h *hub.Hub) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Ошибка WebSocket:", err)
		return
	}
	client := hub.NewClient(roomID, userID, ws, h)
	h.RegisterClient(client)

	go client.Write()
	go client.Read()
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

	wsHub := hub.NewHub(h.commentService, h.commentQuery)
	go wsHub.Run()

	router.GET("/ws/:roomId/:userId", func(c *gin.Context) {
		roomIdStr := c.Param("roomId")
		userIdStr := c.Param("userId")

		roomId, err := strconv.ParseUint(roomIdStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid roomId"})
			return
		}

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
			return
		}

		ServeWS(c, uint(roomId), uint(userId), wsHub)
	})

	return router
}
