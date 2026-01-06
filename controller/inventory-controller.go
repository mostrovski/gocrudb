package controller

import (
	"gocrudb/dto"
	"gocrudb/exception"
	"gocrudb/repository"
	"gocrudb/resource"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			er := ErrorResponse(exception.InvalidRequest{Reason: err.Error()})
			c.JSON(er.StatusCode, er.Body)
			return
		}

		items, err := ic.store.Get(queryDTO.ToQueryConditions())
		if err != nil {
			er := ErrorResponse(err)
			c.JSON(er.StatusCode, er.Body)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}

func (ic InventoryController) Show() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			er := ErrorResponse(exception.InvalidRequest{Reason: "Invalid item id"})
			c.JSON(er.StatusCode, er.Body)
			return
		}

		item, err := ic.store.Find(id)
		if err != nil {
			er := ErrorResponse(err)
			c.JSON(er.StatusCode, er.Body)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": item})
	}
}

func (ic InventoryController) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemDTO dto.CreateItem
		if err := c.ShouldBindJSON(&itemDTO); err != nil {
			er := ErrorResponse(exception.InvalidPayload{Reason: err.Error()})
			c.JSON(er.StatusCode, er.Body)
			return
		}

		instance, ok := resource.Item{}.FromRequestDto(itemDTO).(resource.Item)
		if !ok {
			er := ErrorResponse(exception.InternalServerError{})
			c.JSON(er.StatusCode, er.Body)
			return
		}

		item, err := ic.store.Create(instance)
		if err != nil {
			er := ErrorResponse(err)
			c.JSON(er.StatusCode, er.Body)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": item})
	}
}

func (ic InventoryController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			er := ErrorResponse(exception.InvalidRequest{Reason: "Invalid item id"})
			c.JSON(er.StatusCode, er.Body)
			return
		}

		var itemDTO dto.UpdateItem
		if err := c.ShouldBindJSON(&itemDTO); err != nil {
			er := ErrorResponse(exception.InvalidPayload{Reason: err.Error()})
			c.JSON(er.StatusCode, er.Body)
			return
		}

		instance, ok := resource.Item{}.FromRequestDto(itemDTO).(resource.Item)
		if !ok {
			er := ErrorResponse(exception.InternalServerError{})
			c.JSON(er.StatusCode, er.Body)
			return
		}
		instance.ID = id

		item, err := ic.store.Update(instance)
		if err != nil {
			er := ErrorResponse(err)
			c.JSON(er.StatusCode, er.Body)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": item})
	}
}

func (ic InventoryController) Destroy() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			er := ErrorResponse(exception.InvalidRequest{Reason: "Invalid item id"})
			c.JSON(er.StatusCode, er.Body)
			return
		}

		err = ic.store.Delete(id)
		if err != nil {
			er := ErrorResponse(err)
			c.JSON(er.StatusCode, er.Body)
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
