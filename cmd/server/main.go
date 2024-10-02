package main

import (
    "github.com/gin-gonic/gin"
    "shiftinpro/internal/handlers"
    "shiftinpro/internal/middleware"
    "shiftinpro/internal/repositories"
    "shiftinpro/internal/services"
    "shiftinpro/utility"
)

func main() {
    db := utility.InitDB()
    defer utility.CloseDB()

    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    addressRepo := repositories.NewAddressRepository(db)
    addressService := services.NewAddressService(addressRepo)
    addressHandler := handlers.NewAddressHandler(addressService)

    bookingRepo := repositories.NewBookingRepository(db)
    bookingService := services.NewBookingService(bookingRepo)
    bookingHandler := handlers.NewBookingHandler(bookingService)

    r := gin.Default()
    apiV1 := r.Group("/api/v1")
    {
        user := apiV1.Group("/user")
        {
            user.POST("/register", userHandler.Register)
            user.POST("/login", userHandler.Login)
            user.POST("/send-otp", userHandler.SendOTP)

            //Authorized route
            user.Use(middleware.AuthMiddleware()) // Apply the Auth middleware
            user.POST("/update-address", userHandler.UpdateAddress)

        }

        address := apiV1.Group("/address")
        {
            //Authorized route
            address.Use(middleware.AuthMiddleware()) // Apply the Auth middleware
            address.POST("/get-address", addressHandler.GetAddress)
        }

        booking := apiV1.Group("/booking")
        {
            booking.GET("/seed/item-checklists", bookingHandler.GetItemChecklists)
        }
    }

    if err := r.Run(":8080"); err != nil {
        panic("Failed to start server: " + err.Error())
    }
}
