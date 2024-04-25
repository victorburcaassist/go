package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func AddItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	newItem.ID = uuid.New().String()
	items = append(items, newItem)

	c.JSON(http.StatusOK, newItem)
}

func UpdateItem(c *gin.Context) {
	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	idx := slices.IndexFunc(items, func(i Item) bool {
		return i.ID == updatedItem.ID
	})
	if idx == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})

		return
	}

	items[idx].Name = updatedItem.Name

	c.JSON(http.StatusOK, updatedItem)
}

func DeleteItem(c *gin.Context) {
	var toDelete Item
	if err := c.ShouldBindJSON(&toDelete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	for i, item := range items {
		if item.ID == toDelete.ID {
			items = append(items[:i], items[i+1:]...)

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
