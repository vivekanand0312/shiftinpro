package main

import (
    "github.com/gin-gonic/gin"
    "shiftinpro/internal/handlers"
    "shiftinpro/internal/repository"
    "shiftinpro/internal/services"
    "shiftinpro/utility"
)

func main() {
    db := utility.InitDB()
    defer utility.CloseDB()

    userRepo := repository.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    r := gin.Default()
    apiV1 := r.Group("/api/v1")
    {
        user := apiV1.Group("/user")
        {
            user.POST("/register", userHandler.Register)
            user.POST("/login", userHandler.Login)
        }
    }

    if err := r.Run(":8080"); err != nil {
        panic("Failed to start server: " + err.Error())
    }
}
