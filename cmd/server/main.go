package main

import (
    "github.com/gin-gonic/gin"
    "tg/internal/handlers"
    "tg/internal/repository"
    "tg/internal/services"
    "tg/utility"
)

func main() {
    // Initialize the database connection
    db := utility.InitDB()
    defer utility.CloseDB()

    // Initialize repository and service
    userRepo := repository.NewUserRepository()
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    // Initialize repository and service with DB connection
    tiffinEnqRepo := repository.NewTiffinEnqRepository(db)
    tiffinEnqService := service.NewTiffinEnqService(tiffinEnqRepo)
    tiffinEnqHandler := handler.NewTiffinEnqHandler(tiffinEnqService)

    r := gin.Default()
    // Group routes under api/v1/tiffinEnq
    apiV1 := r.Group("/api/v1")
    {
        tiffinEnq := apiV1.Group("/tiffinEnq") // Use tiffinEnq for the tiffinEnq group
        {
            tiffinEnq.GET("/users/:id", userHandler.GetUser)

            // Add more routes here as needed
            tiffinEnq.POST("/create", tiffinEnqHandler.CreateEnquiry)
            tiffinEnq.GET("/enquiries/:id", tiffinEnqHandler.GetEnquiry)
        }

        user := apiV1.Group("/user")
        {
            user.POST("/register", userHandler.Register)
            // Add more routes here as needed
            // user.POST("/login", userHandler.Login)
        }
    }

    // Start the server
    r.Run(":8080")
}
