package handler

import (
	"fmt"
	approveservice "github.com/TemaStatham/TaskService/client/pkg/app/approve/service"
	commentquery "github.com/TemaStatham/TaskService/client/pkg/app/comment/query"
	commentservice "github.com/TemaStatham/TaskService/client/pkg/app/comment/service"
	responsequery "github.com/TemaStatham/TaskService/client/pkg/app/response/query"
	responseservice "github.com/TemaStatham/TaskService/client/pkg/app/response/service"
	taskquery "github.com/TemaStatham/TaskService/client/pkg/app/task/query"
	taskservice "github.com/TemaStatham/TaskService/client/pkg/app/task/service"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/middleware/auth"
	hub2 "github.com/TemaStatham/TaskService/client/pkg/infrastructure/transport/handler/hub"
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
	taskuserService taskservice.TaskUserServiceInterface
	takuserQuery    taskquery.TaskUserQueryInterface
}

func NewTaskHandler(
	responseQuery responsequery.ResponseQueryInterface,
	responseService responseservice.ResponseServiceInterface,
	commentQuery commentquery.CommentQueryInterface,
	commentService commentservice.CommentServiceInterface,
	taskQuery taskquery.TaskQueryInterface,
	taskService taskservice.TaskServiceInterface,
	approveService approveservice.ApproveServiceInterface,
	taskuserService taskservice.TaskUserServiceInterface,
	takuserQuery taskquery.TaskUserQueryInterface,
) *Handler {
	return &Handler{
		responseQuery,
		responseService,
		commentQuery,
		commentService,
		taskQuery,
		taskService,
		approveService,
		taskuserService,
		takuserQuery,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(ctx *gin.Context, roomID uint, h *hub2.Hub) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Ошибка WebSocket:", err)
		return
	}
	client := hub2.NewClient(1, roomID, ws, h)
	h.RegisterClient(client)

	go client.Write()
	go client.Read()
}

func (h *Handler) Init(jwtSecret string) *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	//todo: переписать парсинг токена на грпс запрос
	//

	httphands := router.Group("/api")
	{
		httphands.Use(auth.UserIdentity(jwtSecret))
		tasksUsers := httphands.Group("/tasks-users")
		{
			tasksUsers.GET("/:id", h.getTasksUsers)
			tasksUsers.POST("/add/:id", h.addTasksUsers)
			tasksUsers.DELETE("/delete/:id", h.deleteTasksUsers)
		}

		tasks := httphands.Group("/tasks")
		{
			tasks.GET("/", h.getTasks)
			tasks.GET("/:id", h.getTask)
			tasks.POST("/", h.createTask)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}

		responses := httphands.Group("/responses")
		{
			responses.GET("/", h.getResponses)
			responses.POST("/", h.createResponse)
			responses.PUT("/:id", h.updateResponse)
		}

		comments := httphands.Group("/comments")
		{
			comments.GET("", h.getComments)
			comments.POST("", h.addComments)
		}

		approves := httphands.Group("/approves")
		{
			approves.POST("", h.addApproves)
		}
	}

	wsHub := hub2.NewHub(h.commentService, h.commentQuery)
	go wsHub.Run()

	router.GET("/ws", func(c *gin.Context) {
		roomIDParam := c.Query("roomID")

		if roomIDParam == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		roomID, err := strconv.ParseUint(roomIDParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		ServeWS(c, uint(roomID), wsHub)
	})

	return router
}
