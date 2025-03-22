package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/middleware/auth"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (h *Handler) getAllByTaskID(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.ShowApproves

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	pag, err := h.approveService.Show(c.Request.Context(), input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"pagination": pag,
	})
}

func (h *Handler) addApproves(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.CreateApprove

	fileBytes, err := base64.StdEncoding.DecodeString(input.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка декодирования файла"})
		return
	}

	// Формируем путь сохранения файла
	saveDir := fmt.Sprintf("uploads/task_%d", input.TaskID) // Папка для конкретного задания
	savePath := filepath.Join(saveDir, fmt.Sprintf("user_%d.png", input.UserID))

	// Создаём директорию, если её нет
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании папки"})
		return
	}

	// Записываем файл
	if err := ioutil.WriteFile(savePath, fileBytes, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	input.File = savePath

	err = h.approveService.Create(c.Request.Context(), input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) rejectApproves(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.RejectApprove

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.approveService.Reject(c.Request.Context(), input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) confirmApproves(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.ConfirmApprove

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.approveService.Confirm(c.Request.Context(), input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
