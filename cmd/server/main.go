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

    addressRepo := repository.NewAddressRepository(db)
    addressService := services.NewAddressService(addressRepo)
    addressHandler := handlers.NewAddressHandler(addressService)

    r := gin.Default()
    apiV1 := r.Group("/api/v1")
    {
        user := apiV1.Group("/user")
        {
            user.POST("/register", userHandler.Register)
            user.POST("/login", userHandler.Login)
            user.POST("/send-otp", userHandler.SendOTP)

            //Auth user
            user.POST("/update-address/:id", userHandler.UpdateAddress)

        }

        address := apiV1.Group("/address")
        {
            address.POST("/get-address", addressHandler.GetAddress)
        }
    }

    if err := r.Run(":8080"); err != nil {
        panic("Failed to start server: " + err.Error())
    }
}
