package main

import (
	"fmt"
	"gocrudb/config"
	"gocrudb/controller"
	"gocrudb/database"
	"gocrudb/middleware"
	"gocrudb/repository"
	"gocrudb/resource"
	"gocrudb/validation"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func main() {
	config.Set()

	database.Setup()
	db := database.Init()
	database.Migrate(db, resource.Item{})
	database.Seed(db, database.GetSeedItems())

	inventoryStore := repository.SqlRepository[uuid.UUID, resource.Item]{}.Init(db)
	inventoryController := controller.InventoryController{}.Init(inventoryStore)

	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for tag, val := range validation.TagValidators() {
			v.RegisterValidation(tag, val)
		}
	}
	router.Use(middleware.RateLimiter())

	router.GET("/inventory", inventoryController.Index())
	router.POST("/inventory", inventoryController.Store())
	router.GET("/inventory/:id", inventoryController.Show())
	router.PATCH("/inventory/:id", inventoryController.Update())
	router.DELETE("/inventory/:id", inventoryController.Destroy())

	router.Run(fmt.Sprintf(":%s", config.Get("app_port")))
}
