package controller

import (
	"errors"
	"gocrudb/dto"
	"gocrudb/repository"
	"gocrudb/resource"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryController struct {
	store       repository.Repository[uuid.UUID, resource.Item]
	initialized bool
}

func (ic InventoryController) Init(r repository.Repository[uuid.UUID, resource.Item]) InventoryController {
	if ic.initialized {
		return ic
	}

	ic.store = r
	ic.initialized = true
	return ic
}

func (ic InventoryController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryDTO dto.QueryItem
		if err := c.ShouldBindQuery(&queryDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		items, err := ic.store.Get(queryDTO.ToQueryConditions())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}

func (ic InventoryController) Show() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item id"})
			return
		}
		item, err := ic.store.Find(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": item})
	}
}

func (ic InventoryController) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemDTO dto.CreateItem
		if err := c.ShouldBindJSON(&itemDTO); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		instance, ok := resource.Item{}.FromReuestDto(itemDTO).(resource.Item)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		item, err := ic.store.Create(instance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": item})
	}
}

func (ic InventoryController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item id"})
			return
		}

		var itemDTO dto.UpdateItem
		if err := c.ShouldBindJSON(&itemDTO); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		instance, ok := resource.Item{}.FromReuestDto(itemDTO).(resource.Item)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		instance.ID = id

		item, err := ic.store.Update(instance)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": item})
	}
}

func (ic InventoryController) Destroy() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item id"})
			return
		}
		err = ic.store.Delete(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
