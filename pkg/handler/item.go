package handler

import (
	root "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomErrorResponse(c, http.StatusBadRequest, "invalid listId param")
		return
	}

	var input root.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newCustomErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoItem.Create(userId, itemId, input)
	if err != nil {
		newCustomErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomErrorResponse(c, http.StatusBadRequest, "invalid listId param")
		return
	}

	items, err := h.service.TodoItem.GetAllItems(userId, itemId)
	if err != nil {
		newCustomErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.service.TodoItem.GetByIdItem(userId, itemId)
	if err != nil {
		newCustomErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input root.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newCustomErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TodoItem.UpdatedByID(userId, id, input)
	if err != nil {
		newCustomErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.TodoItem.DeleteByID(userId, itemId)
	if err != nil {
		newCustomErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
