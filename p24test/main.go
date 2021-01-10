package main

import (
	"p24test/app/controller"
	"p24test/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/api/v1/account/add", controller.CreateAccount) //tanpa auth
	router.POST("/api/v1/login", controller.Login)
	router.GET("/api/v1/account", middleware.Auth, controller.GetAccount) //memeriksa info detil akun yg sedang login
	router.POST("/api/v1/transfer", middleware.Auth, controller.Transfer)
	router.POST("/api/v1/withdraw", middleware.Auth, controller.Withdraw)
	router.POST("/api/v1/deposit", middleware.Auth, controller.Deposit)

	// router.POST("/api/v1/kasihbunga", middleware.Auth, controller.KasihBunga) //function ini NANTINYA hanya boleh dipanggil oleh admin, 1x panggil tiap bulan, untuk semua akun

	router.Run(":8080")
}
