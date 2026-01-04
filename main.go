package main

import (
	"errors"
	"gocrudb/config"
	"gocrudb/database"
	"gocrudb/dto"
	"gocrudb/repository"
	"gocrudb/resource"
	"gocrudb/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = database.Init()
	database.Migrate(db, resource.Item{})
	database.Seed(db, database.GetSeedItems())

	r := repository.SqlRepository[uuid.UUID, resource.Item]{}.Init(db)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for tag, val := range validation.MapCustomByTag() {
			v.RegisterValidation(tag, val)
		}
	}

	router := gin.Default()
	router.GET("/inventory", func(c *gin.Context) {
		var queryDTO dto.QueryItem
		if err := c.ShouldBindQuery(&queryDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		items, err := r.Get(queryDTO.ToQuerySorts())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": items})
	})
	router.POST("/inventory", func(c *gin.Context) {
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

		item, err := r.Create(instance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": item})
	})
	router.GET("/inventory/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item id"})
			return
		}
		item, err := r.Find(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": item})
	})
	router.PATCH("/inventory/:id", func(c *gin.Context) {
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

		item, err := r.Update(instance)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": item})
	})
	router.DELETE("/inventory/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item id"})
			return
		}
		err = r.Delete(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})
	router.Run(config.AppPort())
}
