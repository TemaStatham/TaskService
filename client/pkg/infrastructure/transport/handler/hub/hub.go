package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TemaStatham/TaskService/client/pkg/app/comment/data"
	commentquery "github.com/TemaStatham/TaskService/client/pkg/app/comment/query"
	commentservice "github.com/TemaStatham/TaskService/client/pkg/app/comment/service"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"
	"sync"
)

type Hub struct {
	clients        map[uint]map[*Client]bool // userId -> {клиенты}
	commentService commentservice.CommentServiceInterface
	commentQuery   commentquery.CommentQueryInterface
	register       chan *Client
	unregister     chan *Client
	broadcast      chan Message
	mu             sync.Mutex
}

type Message struct {
	Type   string `json:"type"`
	TaskID uint   `json:"task_id"`
	Data   string `json:"data"`
	UserID uint   `json:"user_id"`
}

func NewHub(
	commentService commentservice.CommentServiceInterface,
	commentQuery commentquery.CommentQueryInterface,
) *Hub {
	return &Hub{
		clients:        make(map[uint]map[*Client]bool),
		commentService: commentService,
		commentQuery:   commentQuery,
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		broadcast:      make(chan Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.RegisterClient(client)
		case client := <-h.unregister:
			h.RemoveClient(client)
		case message := <-h.broadcast:
			h.HandleMessage(message)
		}
	}
}

func (h *Hub) RegisterClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.clients[client.RoomID] == nil {
		h.clients[client.RoomID] = make(map[*Client]bool)
	}
	h.clients[client.RoomID][client] = true
	fmt.Println("Клиент подключился:", client.RoomID)

	comments, err := h.commentQuery.Show(
		context.Background(),
		data.ShowComment{TaskID: client.RoomID, Pagination: paginate.Pagination{}},
	)
	if err == nil {
		client.send <- Message{
			Type:   "history",
			TaskID: client.RoomID,
			Data:   serializeComments(comments),
		}
	}
}

func (h *Hub) RemoveClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.clients[client.RoomID]; ok {
		delete(h.clients[client.RoomID], client)
		close(client.send)
		fmt.Println("Клиент отключён:", client.RoomID)
	}
}

func (h *Hub) HandleMessage(message Message) {
	h.mu.Lock()
	defer h.mu.Unlock()

	switch message.Type {
	case "Создать":
		{
			_, err := h.commentService.Create(
				context.Background(),
				data.CreateComment{Comment: message.Data, TaskID: message.TaskID},
				message.UserID,
			)
			if err != nil {
				fmt.Println("Ошибка сохранения комментария:", err)
				return
			}
		}
	case "Получить":
		{
			var pagination paginate.Pagination
			err := json.Unmarshal([]byte(message.Data), &pagination)
			comments, err := h.commentQuery.Show(
				context.Background(),
				data.ShowComment{TaskID: message.TaskID, Pagination: pagination},
			)
			if err != nil {
				fmt.Println("Ошибка загрузки комментариев:", err)
				return
			}

			h.sendMessageToClients(message.TaskID, comments)
		}
	default:
		return
	}
}

func (h *Hub) sendMessageToClients(taskID uint, comments *paginate.Pagination) {
	message := Message{
		Type:   "message",
		TaskID: taskID,
		Data:   serializeComments(comments),
	}

	for client := range h.clients[taskID] {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients[taskID], client)
		}
	}
}

func serializeComments(comments *paginate.Pagination) string {
	jsonData, err := json.Marshal(comments)
	if err != nil {
		fmt.Println("Ошибка сериализация:", err)
		return "[]"
	}
	return string(jsonData)
}
