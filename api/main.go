package main

import (
	"github.com/jasperwreed/goliath/pkg/handlers"
	"github.com/jasperwreed/goliath/pkg/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Instantiate services
	userService := services.NewUserService(/*db instance*/)
	conversationService := services.NewConversationService(/*db instance*/)
	messageService := services.NewMessageService(/*db instance*/)

	// Instantiate handlers
	userHandler := handlers.NewUserHandler(userService)
	conversationHandler := handlers.NewConversationHandler(conversationService)
	messageHandler := handlers.NewMessageHandler(messageService)

	// User routes
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	// Conversation routes
	conversationGroup := r.Group("/conversations")
	{
		conversationGroup.POST("/", conversationHandler.CreateConversation)
		conversationGroup.GET("/:id", conversationHandler.GetConversationByID)
		conversationGroup.PUT("/:id", conversationHandler.UpdateConversation)
		conversationGroup.DELETE("/:id", conversationHandler.DeleteConversation)
	}

	// Message routes
	messageGroup := r.Group("/messages")
	{
		messageGroup.POST("/", messageHandler.CreateMessage)
		messageGroup.GET("/:id", messageHandler.GetMessageByID)
		messageGroup.GET("/conversation/:id", messageHandler.GetMessagesByConversationID)
		messageGroup.DELETE("/:id", messageHandler.DeleteMessage)
	}

	// Start the Gin server
	r.Run(":8080")
}
